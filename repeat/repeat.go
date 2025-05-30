package iteration

import "strings"

func Repeat(character string, num int) string {
	var repeated strings.Builder

	for i := 0; i < num; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}
