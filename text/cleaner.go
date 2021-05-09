package text

import "unicode"

func Clean(s string) string {
	var result []rune
	for _, r := range s {
		switch {
		case unicode.IsUpper(r):
			result = append(result, unicode.ToLower(r))
		case unicode.IsLower(r):
			result = append(result, r)
		}
	}
	return string(result)
}

