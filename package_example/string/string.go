package string

import "strings"

func Reverse(s string) string {
	b := []rune(s)

	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

type groupChars []rune

func VowelsOnly(s string) string {
	s = strings.ToLower(s)
	new_string := ""

	for _, value := range s {
		switch value {
		case 'a', 'e', 'i', 'o', 'u':
			new_string = new_string + string(value)
		}
	}
	return string(new_string)
}
