package main

type Language string

const (
	Spanish Language = "es"
	English Language = "en"
	German  Language = "de"
	Italian Language = "it"
	French  Language = "fr"
)

func Parse(input string) (Language, bool) {
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
