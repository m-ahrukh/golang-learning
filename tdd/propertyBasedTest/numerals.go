package propertybasedtest

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var RomanNumerals = []RomanNumeral{
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
