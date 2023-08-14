package main

import (
  "net/http"
  "fmt"
)

func handleName(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  
  name := r.URL.Path
  fmt.Fprintf(w, "Hello, %s!\n", name)
}

func hanldeNoName(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)

  fmt.Fprintf(w, "Hello, World!")
}

func main() {
  mux := http.NewServeMux()
  mux.Handle("/hello/", http.StripPrefix("/hello/", http.HandlerFunc(handleName)))
  
  server := http.Server {
    Addr: ":8080",
    Handler: mux,
  }

  if err := server.ListenAndServe(); err != nil {
    panic(err)
  }
}
