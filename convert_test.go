package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRomanNumeralAcceptsValidNumerals(t *testing.T) {
	n := "VII"
	rn, err := NewRomanNumeral(n)

	if err != nil {
		t.Fatal(fmt.Sprintf("NewRomanNumeral returned an error: %v", err))
	}

	if rn.String() != "VII" {
		t.Fatal(fmt.Sprintf("NewRomanNumeral return with unxepected numeral value expected: %v actual: %v", n, rn))
	}
}

func TestNewRomanNumeralConvertsRomanNumeralToCorrectCase(t *testing.T) {
	rn, err := NewRomanNumeral("vii")

	if err != nil {
		t.Fatal(fmt.Sprintf("NewRomanNumeral returned an error: %v", err))
	}

	if rn.String() != "VII" {
		t.Fatal(fmt.Sprintf("NewRomanNumeral return with unxepected numeral value expected: VII actual: %v", rn))
	}
}

func TestNewRomanNumeralDoesNotAcceptInvalidCharacters(t *testing.T) {
	rn, err := NewRomanNumeral("viia")

	if err == nil {
		t.Fatal("NewRomanNumeral did not return an error when one was expected")
	}

	if rn.String() != "" {
		t.Fatal(fmt.Sprintf("NewRomanNumeral return with a RomanNumeral with value %v empty string expected", rn))
	}
}

func TestReverseReversesTheOrderOfTheASlice(t *testing.T) {
	givenSlice := []string{"abc", "def", "ghi"}
	expectedSlice := []string{"ghi", "def", "abc"}
	reversedSlice := reverse(givenSlice)

	assert.Equal(t, expectedSlice, reversedSlice)
}

func TestGetNumeralValueForIis1(t *testing.T) {
	assert.Equal(t, 1, getNumeralValue("I"))
}

func TestGetNumeralValueForVis5(t *testing.T) {
	assert.Equal(t, 5, getNumeralValue("V"))
}

func TestGetNumeralValueForXis10(t *testing.T) {
	assert.Equal(t, 10, getNumeralValue("X"))
}

func TestGetNumeralValueForLis50(t *testing.T) {
	assert.Equal(t, 50, getNumeralValue("L"))
}

func TestGetNumeralValueForCis100(t *testing.T) {
	assert.Equal(t, 100, getNumeralValue("C"))
}

func TestGetNumeralValueForDis500(t *testing.T) {
	assert.Equal(t, 500, getNumeralValue("D"))
}

func TestGetNumeralValueForMis1000(t *testing.T) {
	assert.Equal(t, 1000, getNumeralValue("M"))
}

func TestConvertToIntegerXVis15(t *testing.T) {
	rn, _ := NewRomanNumeral("XV")

	assert.Equal(t, 15, convertToInteger(rn))
}

func TestConvertToIntegerXIVis14(t *testing.T) {
	rn, _ := NewRomanNumeral("XIV")

	assert.Equal(t, 14, convertToInteger(rn))
}

func TestConvertToIntegerMCMXCVIIIis1998(t *testing.T) {
	rn, _ := NewRomanNumeral("MCMXCVIII")

	assert.Equal(t, 1998, convertToInteger(rn))
}
