package main

import (
  "fmt"
  "net/http"
)

func main() {
  handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Hello, World!")
  })

  server := http.Server {
    Addr: ":8080",
    Handler: handler,
  }

  if err := server.ListenAndServe(); err != nil {
    panic(err)
  }

}
