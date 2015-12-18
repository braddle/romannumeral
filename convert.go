package romannumeral

import (
	"errors"
	"regexp"
	"strings"
)

type RomanNumeral string

func (r RomanNumeral) String() string {
	return string(r)
}

func (r RomanNumeral) asSlice() []string {
	return strings.Split(string(r), "")
}

func (r RomanNumeral) ToInt() int {
	total := 0
	var previous int

	numerals := r.asSlice()

	for i := len(numerals) - 1; i >= 0; i-- {
		numeral := numerals[i]
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

func getNumeralValue(n string) int {
	switch n {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}

func NewRomanNumeral(numeral string) (RomanNumeral, error) {
	numeral = strings.ToUpper(strings.Join(strings.Fields(numeral), ""))

	reg := regexp.MustCompile("[^I|V|X|L|C|D|M]")
	invalid := reg.MatchString(numeral)

	if invalid {
		return RomanNumeral(""), errors.New("Invalid characters in roman numeral. Valid roman numerals are I, V, X, L, C, D and M")
	}

	return RomanNumeral(numeral), nil
}
