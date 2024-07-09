package main

func normalize(phone string) string {

	//normaize number by iterating string
	//normalize number using regex

	//Method 1.a
	output := ""
	for _, ch := range phone {
		if ch >= '0' && ch <= '9' {
			output = output + string(ch)
		}
	}
	return output

	//Method 1.b
	// var buf bytes.Buffer
	// for _, ch := range phone {
	// 	if ch >= '0' && ch <= '9' {
	// 		buf.WriteRune(ch)
	// 	}
	// }
	// return buf.String()

	//Method 2
	// re := regexp.MustCompile("[0-9]+")
	// matches := re.FindAllString(phone, -1)
	// fmt.Println(strings.Join(matches, ""))
	// return strings.Join(matches, "")
}
