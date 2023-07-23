package main

import (
	"fmt"
	"os"
	"strings"
)

func reorderName(firstName, lastName, countryCode, surname string) string {
	switch strings.ToUpper(countryCode) {
	case "US", "AUS":
		return fmt.Sprintf("%s %s %s", lastName, surname, firstName)
	case "CN", "JP":
		return fmt.Sprintf("%s %s", firstName, lastName)
	case "VN":
		return fmt.Sprintf("%s %s %s", firstName, surname, lastName)
	default:
		return fmt.Sprintf("%s %s %s", firstName, lastName, surname)
	}
}

func main() {
	// a slice of strings that contains the command-line arguments passed to the program
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("Usage: go run . firstName lastName surname countryCode")
		fmt.Println("surname is optional")
		return
	}

	countryCode := ""
	surname := ""
	if len(args) > 3 {
		countryCode = args[3]
		surname = args[2]
	} else {
		countryCode = args[2]
	}

	result := reorderName(args[0], args[1], countryCode, surname)
	fmt.Println("output: ", result)
}
