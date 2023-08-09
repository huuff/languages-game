package main

import (
	"fmt"
	"os"

	"hello/level3/language"
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
	greeters := map[language.Language]greeterFunction{
		language.Spanish: func(name string) string {
			return fmt.Sprintf("¡Hola, %s!", defaultString(name, "Mundo"))
		},
		language.English: func(name string) string {
			return fmt.Sprintf("Hello, %s!", defaultString(name, "World"))
		},
		language.French: func(name string) string {
			return fmt.Sprintf("Bonjour, %s!", defaultString(name, "le Monde"))
		},
		language.Italian: func(name string) string {
			return fmt.Sprintf("Ciao, %s!", defaultString(name, "Mondo"))
		},
		language.German: func(name string) string {
			return fmt.Sprintf("Hallo, %s!", defaultString(name, "Welt"))
		},
	}

	var name string
	var lang language.Language

	if len(os.Args) > 1 {
		name = os.Args[1]
	} else {
		name = ""
	}
	if len(os.Args) > 2 {
		lang, _ = language.Parse(os.Args[2])
	} else {
		lang = language.English
	}

	greet, ok := greeters[lang]
	if ok {
		fmt.Println(greet(name))
	} else {
		fmt.Println(fmt.Sprintf("Unrecognized language %s", lang))
	}
}
