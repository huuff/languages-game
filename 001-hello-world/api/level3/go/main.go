package main

import (
  "net/http"
  "fmt"
)

func getSalutation(name string, language Language) string {
  return greeters[language](name)
}

func handleHelloName(w http.ResponseWriter, r *http.Request) {
  name := r.URL.Path
  // TODO: Put the language extraction stuff into a different function
  language := defaultString(r.Header.Get("accept-language"), "en")
  actualLanguage, ok := Parse(language)
  if !ok {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "Unrecognized language %s", actualLanguage)
    return
  }

  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "%s\n", getSalutation(name, actualLanguage))
}

func main() {
  mux := http.NewServeMux()
  mux.Handle("/hello/", http.StripPrefix("/hello/", http.HandlerFunc(handleHelloName)))

  server := http.Server {
    Addr: ":8080",
    Handler: mux,
  }

  if err := server.ListenAndServe(); err != nil {
    panic(err)
  }

}
