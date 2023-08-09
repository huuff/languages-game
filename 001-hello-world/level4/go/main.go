package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

type greeterFunction func(name string) string 

func defaultString(input, defaultValue string) string {
  if input != "" {
    return input
  } else {
    return defaultValue
  }
}

// TODO: Put it in another file
type Language string
const (
  Spanish Language = "es"
  English Language = "en"
  German Language = "de"
  Italian Language = "it"
  French Language = "fr"
)
type greetersMap map[Language]greeterFunction

func toLanguage(input string) (Language, bool) {
  switch input {
  case "es":
    return Spanish, true
  case "en":
    return English, true
  case "de":
    return German, true
  case "it":
    return Italian, true
  case "fr":
    return French, true
  default:
    return "", false
  }
}

func newHelloHandler(greeters greetersMap) func(w http.ResponseWriter, req *http.Request) {
  return func(w http.ResponseWriter, req *http.Request) {
    headerLanguage := req.Header.Get("Accept-Language")
    var language Language
    if headerLanguage != "" {
      language, _ = toLanguage(headerLanguage) 
    } else {
      language = English
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
  greeters := map[Language] greeterFunction {
    Spanish: func(name string) string { 
      return fmt.Sprintf("¡Hola, %s!", defaultString(name, "Mundo"))
    },
    English: func(name string) string { 
      return fmt.Sprintf("Hello, %s!", defaultString(name, "World"))
    },
    French: func(name string) string {
      return fmt.Sprintf("Bonjour, %s!", defaultString(name, "le Monde"))
    },
    Italian: func(name string) string { 
      return fmt.Sprintf("Ciao, %s!", defaultString(name, "Mondo"))
    },
    German: func(name string) string {
      return fmt.Sprintf("Hallo, %s!", defaultString(name, "Welt"))
    },
  }

  router := mux.NewRouter()
  router.HandleFunc("/hello/{name}", newHelloHandler(greeters))
  http.ListenAndServe(":3333", router)
}
