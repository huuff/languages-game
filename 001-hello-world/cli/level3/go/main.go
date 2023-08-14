package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	var lang Language

	if len(os.Args) > 1 {
		name = os.Args[1]
	} else {
		name = ""
	}
	if len(os.Args) > 2 {
		lang, _ = Parse(os.Args[2])
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
