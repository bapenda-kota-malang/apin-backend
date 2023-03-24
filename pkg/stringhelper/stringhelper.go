package stringhelper

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// alt string.title next akan depreceted
func UpperCaseWord(s string) string {
	s = strings.ToLower(s)
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			if unicode.IsLower(r) {
				return unicode.ToUpper(r)
			}
			return r
		}
		return -1
	}, s)
}
func IsAlphanumericAllSymbol(s string) string {
	re := regexp.MustCompile(`[^[:alnum:][:punct:]]`)
	return re.ReplaceAllString(s, "")
}

func IsNumeric(s string) string {
	re := regexp.MustCompile(`[^0-9]`)
	return re.ReplaceAllString(s, "")
}

func IsAlphabetic(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z]`)
	return re.ReplaceAllString(s, "")
}

func NullToAny(s *string, anyVal string) string {
	if s == nil {
		return anyVal
	}
	if *s == "" {
		return anyVal
	}
	return *s
}

func BoolToAny(s *bool, anyTrueVal string, anyFalseVal string) string {
	if s == nil {
		return anyFalseVal
	}
	if *s {
		return anyTrueVal
	}
	return anyFalseVal
}
func FloatToAny(s *float64, anyVal string) string {
	if s == nil {
		return anyVal
	}
	return strconv.FormatFloat(*s, 'f', 2, 64)
}
