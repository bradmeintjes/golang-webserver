package routing

import "net/http"

// Get will wrap the given handler in a request method assertion
func Get(next http.HandlerFunc) http.HandlerFunc {
	return Method(http.MethodGet, next)
}

// Put will wrap the given handler in a request method assertion
func Put(next http.HandlerFunc) http.HandlerFunc {
	return Method(http.MethodPut, next)
}

// Post will wrap the given handler in a request method assertion
func Post(next http.HandlerFunc) http.HandlerFunc {
	return Method(http.MethodPut, next)
}

// Delete will wrap the given handler in a request method assertion
func Delete(next http.HandlerFunc) http.HandlerFunc {
	return Method(http.MethodDelete, next)
}

// Method will wrap the given handler in a method check to ensure only
// requests with the given method are handled
func Method(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			next(w, r)
		} else {
			illegalMethod(w)
		}
	}
}

func illegalMethod(w http.ResponseWriter) {
	http.Error(w, "illegal method", http.StatusNotFound)
}
