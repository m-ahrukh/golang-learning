package iteration

var repeatedCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatedCount; i++ {
		repeated = repeated + character
	}
	return repeated
}

func RepeatChar(character string, times int) string {
	if times < 0 {
		return "Number of times must be positive"
	}
	var repeated string
	for i := 0; i < times; i++ {
		repeated = repeated + character
	}
	return repeated
}
