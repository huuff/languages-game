package main

func defaultString(input, defaultValue string) string {
	if input != "" {
		return input
	} else {
		return defaultValue
	}
}
