# Roman Numeral

First attempt at programming in Go

## PACKAGE DOCUMENTATION

package romannumeral
    import "romannumeral"


### TYPES

type RomanNumeral string

func NewRomanNumeral(numeral string) (RomanNumeral, error)
    NewRomanNumeral creates an instance of a RomanNumeral from the given
        numeral string. If the numeral contains invalid characters an error is
	    returned

func (r RomanNumeral) String() string
    String returns the RomanNumeral as a string

func (r RomanNumeral) ToInt() int
    ToInt returns an int representation of the RomanNumeral