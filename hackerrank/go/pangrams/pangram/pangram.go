package pangram

import "strings"

func IsPangram(s string) string {
	seenLetters := make(map[rune]bool)
	lowerCase := strings.ToLower(s)

	for _, char := range lowerCase {
		if char >= 'a' && char <= 'z' {
			seenLetters[char] = true
		}
	}

	if len(seenLetters) == 26 {
		return "pangram"
	}
	return "not pangram"
}
