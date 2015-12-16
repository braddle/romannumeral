package main

import (
    "fmt"
    "os"
    "strings"
)

type RomanNumeral struct {
  numeral string
}

func  NewRomanNumeral(numeral string) RomanNumeral  {
  rn := RomanNumeral { numeral }
  // rn.numeral = numeral

  return rn
}

func (r RomanNumeral) String() string {
  return r.numeral
}

func (r RomanNumeral) asSlice() []string {
  return strings.Split(r.numeral, "")
}

func main() {
  numeralString := os.Args[1]
  numeral := NewRomanNumeral(numeralString)
  number := convertToInteger(numeral)

  fmt.Printf("%v = %v\n", numeral, number)
}

func convertToInteger(rn RomanNumeral) int {
  total := 0
  var previous int

  numerals := reverse(rn.asSlice())

  for _, numeral := range numerals {
    value := getNumeralValue(numeral)

    fmt.Printf("%v\n", previous)
    if (previous <= value) {
      total = total + value
    } else {
      total = total - value
    }

    previous = value

  }

  // fmt.Printf("%v", numerals)
  return total
}

func reverse(sl []string) []string {
  for i := len(sl)/2-1; i >= 0; i-- {
    opp := len(sl)-1-i
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
