package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	app := application{
		greeters: greeters,
	}

	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}", app.helloHandler)
	http.ListenAndServe(":3333", router)
}

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
