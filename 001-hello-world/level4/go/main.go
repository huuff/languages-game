package main

import (
	"fmt"
	"hello/level4/language"
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

type greetersMap map[language.Language]greeterFunction

func newHelloHandler(greeters greetersMap) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		headerLanguage := req.Header.Get("Accept-Language")
		var actualLanguage language.Language
		if headerLanguage != "" {
			actualLanguage, _ = language.Parse(headerLanguage)
		} else {
			actualLanguage = language.English
		}

		greeter, ok := greeters[actualLanguage]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Unrecognized language %s\n", actualLanguage)
			return
		}

		pathVars := mux.Vars(req)
		name, _ := pathVars["name"]

		fmt.Fprintf(w, greeter(name)+"\n")
	}
}

func main() {
	greeters := greetersMap{
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

	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}", newHelloHandler(greeters))
	http.ListenAndServe(":3333", router)
}
