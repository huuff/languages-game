package main

import (
  "fmt"
  "os"
)

type greeterFunction func(name string) string 

func main() {
  greeters := map[string] greeterFunction {
    "es": func(name string) string { return fmt.Sprintf("Hola mundo")},
    "en": func(name string) string { return fmt.Sprintf("Hello world")},
    "fr": func(name string) string { return fmt.Sprintf("Bonjour le monde")},
    "it": func(name string) string { return fmt.Sprintf("AAAAAAAAAAAa")},
  }

  if len(os.Args) > 1 {
    lang := os.Args[1]
    result, ok := greeters[lang]
    if ok {
      fmt.Println(result("whatev"))
    } else {
      fmt.Println(fmt.Sprintf("Unrecognized language %s", lang))
    }
  } else {
    fmt.Println("No lang!")
  }
}
