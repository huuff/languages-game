package main

import (
  "fmt"
  "os"
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

  var name string
  var lang Language

  if len(os.Args) > 1 {
    name = os.Args[1]
  } else {
    name = ""
  }
  if len(os.Args) > 2 {
    lang, _ = toLanguage(os.Args[2])
  } else {
    lang = English
  }

  greet, ok := greeters[lang]
  if ok {
    fmt.Println(greet(name))
  } else {
    fmt.Println(fmt.Sprintf("Unrecognized language %s", lang))
  }
}
