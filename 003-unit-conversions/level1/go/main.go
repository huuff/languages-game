package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const usageMessage = "Wrong arguments. Use <amount><source unit> to <target unit>.\n"

// TODO: Should actually parse a float
func parseNumber(input string) int {
	re := regexp.MustCompile(`^\d+`)
	matches := re.FindStringSubmatch(input)
	if len(matches) != 1 {
		panic(fmt.Sprintf("Can't parse the number in `%s`. Expecting 1 match but got `%d`", input, len(matches)))
	}

	value, err := strconv.Atoi(matches[0])
	if err != nil {
		panic(fmt.Sprintf("Error converting to string: `%s`", matches[0]))
	}

	return value
}

func main() {
	if (len(os.Args) != 4) || (os.Args[2] != "to") {
		panic(usageMessage)
		os.Exit(1)
	}

	source := os.Args[1]

	// TODO: Validate
	target := os.Args[3]

	sourceAmount := parseNumber(source)

	fmt.Fprintf(os.Stdout, "Source amount: %d\n", sourceAmount)
	fmt.Fprintf(os.Stdout, "Target unit: %s\n", target)
}
