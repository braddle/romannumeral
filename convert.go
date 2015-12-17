package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type RomanNumeral struct {
	numeral string
}

func NewRomanNumeral(numeral string) (RomanNumeral, error) {
	numeral = strings.ToUpper(strings.Join(strings.Fields(numeral), ""))

	invalid, err := regexp.MatchString("[^I|V|X|L|C|D|M]", numeral)

	if err != nil {
		log.Fatal(err)
	}

	if invalid {
		return RomanNumeral{}, errors.New("Invalid characters in roman numeral. Valid roman numerals are I, V, X, L, C, D and M")
	}

	rn := RomanNumeral{numeral}

	return rn, nil
}

func (r RomanNumeral) String() string {
	return r.numeral
}

func (r RomanNumeral) asSlice() []string {
	return strings.Split(r.numeral, "")
}

func main() {
	numeralString := os.Args[1]
	numeral, err := NewRomanNumeral(numeralString)

	if err != nil {
		log.Fatal(err)
	}

	number := convertToInteger(numeral)

	fmt.Printf("%v = %v\n", numeral, number)
}

func convertToInteger(rn RomanNumeral) int {
	total := 0
	var previous int

	numerals := reverse(rn.asSlice())

	for _, numeral := range numerals {
		value := getNumeralValue(numeral)

		if previous <= value {
			total = total + value
		} else {
			total = total - value
		}

		previous = value
	}

	return total
}

func reverse(sl []string) []string {
	for i := len(sl)/2 - 1; i >= 0; i-- {
		opp := len(sl) - 1 - i
		sl[i], sl[opp] = sl[opp], sl[i]
	}

	return sl
}

func getNumeralValue(n string) int {
	numeralValues := make(map[string]int)
	numeralValues["I"] = 1
	numeralValues["V"] = 5
	numeralValues["X"] = 10
	numeralValues["L"] = 50
	numeralValues["C"] = 100
	numeralValues["D"] = 500
	numeralValues["M"] = 1000

	return numeralValues[n]
}
