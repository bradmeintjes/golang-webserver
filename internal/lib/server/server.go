package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type key int

const (
	requestIDKey key = 0
)

// New will create a new http.Server instance with a secure configuration
func New(logger *log.Logger, mux *http.ServeMux, addr string) *http.Server {
	// minimum tls configuration for a secure golang web server
	// https://blog.cloudflare.com/exposing-go-on-the-internet/
	tlsConfig := &tls.Config{
		// Causes servers to use Go's default ciphersuite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519, // Go 1.8 only
		},
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,

			// Best disabled, as they don't provide Forward Secrecy,
			// but might be necessary for some clients
			// tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			// tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	return &http.Server{
		Addr:         addr,
		Handler:      tracing(nextRequestID)(logging(logger)(mux)),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
	}
}

func nextRequestID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Serve will wrap the call to ListenAndServe in a graceful shutdown handler
func Serve(logger *log.Logger, svr *http.Server) error {
	return gracefully(logger, svr, func(s *http.Server) error {
		return s.ListenAndServe()
	})
}

// ServeTLS will wrap the call to ListenAndServeTLS in a graceful shutdown handler
func ServeTLS(logger *log.Logger, svr *http.Server, certFile, certKey string) error {
	return gracefully(logger, svr, func(s *http.Server) error {
		return s.ListenAndServeTLS(certFile, certKey)
	})
}

// spawns a routine to wait for a interrupt signal and handle the shutdown gracefully
func gracefully(logger *log.Logger, svr *http.Server, serve func(*http.Server) error) error {
	done := make(chan error)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		svr.SetKeepAlivesEnabled(false)
		if err := svr.Shutdown(ctx); err != nil {
			done <- fmt.Errorf("could not gracefully shutdown the server: %s", err)
		}
		done <- nil
	}()

	err := serve(svr)

	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("could not listen on %s: %v", svr.Addr, err)
	}

	return <-done
}
