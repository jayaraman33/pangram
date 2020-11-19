package pangram

import "unicode"

func IsPangram(s string) bool {

	m := make(map[rune]bool)

	for _, l := range s {

		if unicode.IsLetter(l) {

			m[unicode.ToLower(l)] = true

		}
	}

	return len(m) == 26
}
