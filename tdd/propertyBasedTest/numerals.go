package propertybasedtest

import (
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

var RomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	// if arabic == 3 {
	// 	return "III"
	// }
	// if arabic == 2 {
	// 	return "II"
	// }
	// return "I"

	var result strings.Builder

	// for i := arabic; i > 0; i-- {
	// 	if i == 5 {
	// 		result.WriteString("V")
	// 		break
	// 	}

	// 	if i == 4 {
	// 		result.WriteString("IV")
	// 		break
	// 	}

	// 	result.WriteString("I")
	// }

	// for arabic > 0 {
	// 	switch {
	// 	case arabic > 8:
	// 		result.WriteString("IX")
	// 		arabic -= 9
	// 	case arabic > 4:
	// 		result.WriteString("V")
	// 		arabic -= 5
	// 	case arabic > 3:
	// 		result.WriteString("IV")
	// 		arabic -= 4
	// 	default:
	// 		result.WriteString("I")
	// 		arabic--
	// 	}
	// }

	for _, numeral := range RomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	// if roman == "III" {
	// 	return 3
	// }
	// if roman == "II" {
	// 	return 2
	// }
	// return 1

	// total := 0
	// for range roman {
	// 	total++
	// }
	// return total

	var arabic = 0
	for _, numeral := range RomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			arabic += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}
	return arabic
}
