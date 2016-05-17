package main

import "net/http"

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
  w.Write([]byte("Hello todo"))
}