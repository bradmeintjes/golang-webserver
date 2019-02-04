package todo

import (
	"encoding/json"
	"go/webserver/internal/common/routing"
	"net/http"
)

const (
	todoIDContextKey = "todo-id"
)

/*
 * ServeHTTP allows this package to implement the http.Handler interface
 * Routing is accomplished through an iterative handling process through
 * each level of the API, each level passing down a shifted path to the
 * appropriate next handler
 *
 * For this class the handler will expect any of the following methods
 *  GET /			retrieves all todos
 * 	GET /{id}		retrieves a specific todo with id
 *	DELETE /{id}	deletes a todo with the given id
 *  PUT /{id}		updates the todo with id
 *	POST /			creates a new todo
 * 	PUT /allDone	updates all todos to complete
 *
 * The path parsing above is done simpy in code. Sure, it may be more verbose
 * than simply typing a route with some wildcards, but it strikes a better balance
 * between speed, simplicity, and std lib compliance, negating the need for
 * complex/heavy/opinionated routing libs
 *
 * This is also a very common structure for handling crud operations, possibly
 * can be abstracted out and have the handlers just implement a simpe crud interface?
 */
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = routing.ShiftPath(r.URL.Path)

	var handler http.HandlerFunc
	switch head {
	case "":
		switch r.Method {
		case http.MethodGet:
			handler = getAll
		case http.MethodPost:
			handler = create
		}

	case "allDone":
		handler = routing.Post(markAllDone)

	default: // assume head is an id
		switch r.Method {
		case http.MethodGet:
			handler = withIDContext(head, get)
		case http.MethodPut:
			handler = withIDContext(head, update)
		case http.MethodDelete:
			handler = withIDContext(head, delete)
		}
	}

	if handler == nil {
		http.Error(w, "route not found", http.StatusNotFound)
		return
	}

	handler(w, r)
}

func getAll(w http.ResponseWriter, r *http.Request) {

}

func get(w http.ResponseWriter, r *http.Request) {
}

func create(w http.ResponseWriter, r *http.Request) {
}

func update(w http.ResponseWriter, r *http.Request) {
}

func delete(w http.ResponseWriter, r *http.Request) {
}

func markAllDone(w http.ResponseWriter, r *http.Request) {

}

func withIDContext(id string, handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), todoIDContextKey, id)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

func asJSON(w http.ResponseWriter, val interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	err := json.NewEncoder(w).Encode(val)
	if err != nil {
		errs.WriteHttp(w, err)
		return
	}
}
