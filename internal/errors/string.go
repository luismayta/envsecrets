package errors

import (
	"unicode"
)

// ToUnderScore converts CamelCase to snake_case
func ToUnderScore(name string) string {
	result := make([]rune, 0, len(name))

	for i, char := range name {
		if unicode.IsUpper(char) && i > 0 {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(char))
	}

	return string(result)
}
