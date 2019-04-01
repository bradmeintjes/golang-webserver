package server

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"syscall"
	"testing"
	"time"
)

func TestCreateServer(t *testing.T) {
	// test does not run on windows since there is no interrupt signal capability
	if runtime.GOOS == "windows" {
		t.Log("skipping server test on windows")
		t.SkipNow()
		return
	}

	logger := log.New(os.Stdout, "test", log.LstdFlags)
	router := http.NewServeMux()

	go func() {
		time.Sleep(2 * time.Second)
		// send a terminate signal - does not work in windows
		p, err := os.FindProcess(syscall.Getpid())
		if err == nil {
			t.Log("killing process")
			err = p.Signal(os.Interrupt)
		}
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()

	svr := New(logger, router, ":9999")
	t.Log("starting server")
	err := Serve(logger, svr)

	if err != nil {
		t.Logf("failed to serve %s\n", err)
		t.FailNow()
	}
}
