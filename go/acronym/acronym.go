package acronym

import (
	"fmt"
	"strings"
	"unicode"
)

const testVersion = 1

func abbreviate(phrase string) string {
	acronym := ""
	for _, char := range strings.Title(strings.Split(phrase, ":")[0]) {
		if unicode.IsUpper(char) {
			acronym = fmt.Sprintf("%v%v", acronym, string(char))
		}
	}
	return acronym
}
