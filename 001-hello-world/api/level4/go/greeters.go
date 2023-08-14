package main

import "fmt"

type greeterFunction func(name string) string
type greetersMap map[Language]greeterFunction

var greeters = greetersMap{
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
