// now we NEED to do `go build` and then run the executable
// `go run` can only run single files and doesn't resolve those linked
package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/todo", todoHandler)

	http.ListenAndServe(":8080", mux)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello server3"))
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	// jin
	// negroni
	// gorilla mux

	switch r.Method {
	case "GET":
		todo := GetAll() // GetTodo is found automatically from todo.go as it is in same package

		result, err := json.Marshal(todo)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)

	case "POST":
		todo := Todo{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&todo)
		if err != nil {
			panic(err)
		}
		SaveTodo(todo)
		w.Write([]byte("POST"))

	case "PUT":
		w.Write([]byte("PUT"))
	case "DELETE":
		w.Write([]byte("DELETE"))
	}

}
