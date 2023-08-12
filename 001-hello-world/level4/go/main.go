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

type greetersMap map[Language]greeterFunction

type application struct {
	greeters greetersMap
}

func (app *application) helloHandler(w http.ResponseWriter, req *http.Request) {
	headerLanguage := req.Header.Get("Accept-Language")
	var actualLanguage Language
	if headerLanguage != "" {
		actualLanguage, _ = Parse(headerLanguage)
	} else {
		actualLanguage = English
	}

	greeter, ok := app.greeters[actualLanguage]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unrecognized language %s\n", actualLanguage)
		return
	}

	pathVars := mux.Vars(req)
	name, _ := pathVars["name"]

	fmt.Fprintf(w, greeter(name)+"\n")
}

func main() {
	greeters := greetersMap{
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

	app := application{
		greeters: greeters,
	}

	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}", app.helloHandler)
	http.ListenAndServe(":3333", router)
}
