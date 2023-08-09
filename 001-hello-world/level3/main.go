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

func main() {
  greeters := map[string] greeterFunction {
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

  var lang, name string

  if len(os.Args) > 1 {
    name = os.Args[1]
  } else {
    name = ""
  }
  if len(os.Args) > 2 {
    lang = os.Args[2]
  } else {
    lang = "en"
  }

  greet, ok := greeters[lang]
  if ok {
    fmt.Println(greet(name))
  } else {
    fmt.Println(fmt.Sprintf("Unrecognized language %s", lang))
  }
}
