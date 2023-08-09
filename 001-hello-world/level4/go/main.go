package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

type greeterFunction func(name string) string 
type greetersMap map[string]greeterFunction

func defaultString(input, defaultValue string) string {
  if input != "" {
    return input
  } else {
    return defaultValue
  }
}

func newHelloHandler(greeters greetersMap) func(w http.ResponseWriter, req *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    language := req.Header.Get("Accept-Language")
    if language == "" {
      language = "en"
    }

    greeter, ok := greeters[language]
    if !ok {
      w.WriteHeader(http.StatusBadRequest)
      fmt.Fprintf(w, "Unrecognized language %s\n", language)
      return
    }

    pathVars := mux.Vars(req)
    name, _ := pathVars["name"]

    fmt.Println(fmt.Sprintf("Language: %s", language))
    fmt.Println(fmt.Sprintf("Name: %s", name))
    fmt.Fprintf(w, greeter(name) + "\n")
  }
}

func main() {
  // TODO: Use an iota for the language values (also in level3)
  greeters := greetersMap {
    "es": func(name string) string { 
      return fmt.Sprintf("¡Hola, %s!", defaultString(name, "Mundo"))
    },
    "en": func(name string) string { 
      return fmt.Sprintf("Hello, %s!", defaultString(name, "World"))
    },
    "fr": func(name string) string {
      return fmt.Sprintf("Bonjour, %s!", defaultString(name, "le Monde"))
    },
    "it": func(name string) string { 
      return fmt.Sprintf("Ciao, %s!", defaultString(name, "Mondo"))
    },
    "de": func(name string) string {
      return fmt.Sprintf("Hallo, %s!", defaultString(name, "Welt"))
    },
  }

  router := mux.NewRouter()
  router.HandleFunc("/hello/{name}", newHelloHandler(greeters))
  http.ListenAndServe(":3333", router)
}
