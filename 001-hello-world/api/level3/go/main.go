package main

import (
  "net/http"
  "fmt"
  "errors"
)

func getSalutation(name string, language Language) string {
  return greeters[language](name)
}

func getLanguage(r *http.Request) (Language, error) {
  language := defaultString(r.Header.Get("accept-language"), "en")
  actualLanguage, ok := Parse(language)
  if !ok {
    return "", errors.New(fmt.Sprintf("Unrecognized language: %s", language))
  } else {
    return actualLanguage, nil
  }
}

func handleHelloNoName(w http.ResponseWriter, r *http.Request) {
  language, err := getLanguage(r)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "%s\n", err.Error())
    return
  }

  w.Header().Set("content-language", string(language))
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "%s\n", getSalutation("", language))
}

func handleHelloName(w http.ResponseWriter, r *http.Request) {
  name := r.URL.Path
  language, err := getLanguage(r)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprintf(w, "%s\n", err.Error())
    return
  }

  w.Header().Set("content-language", string(language))
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "%s\n", getSalutation(name, language))
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/hello", handleHelloNoName)
  mux.Handle("/hello/", http.StripPrefix("/hello/", http.HandlerFunc(handleHelloName)))

  server := http.Server {
    Addr: ":8080",
    Handler: mux,
  }

  if err := server.ListenAndServe(); err != nil {
    panic(err)
  }

}
