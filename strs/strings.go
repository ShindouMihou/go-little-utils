package strs

import "strings"

// CaseInsensitiveHasPrefix checks whether the source string has a prefix that matches case-insensitively.
func CaseInsensitiveHasPrefix(source, prefix string) bool {
	return len(source) >= len(prefix) && strings.EqualFold(source[0:len(prefix)], prefix)
}

// CaseInsensitiveHasSuffix checks whether the source string has a suffix that matches case-insensitively.
func CaseInsensitiveHasSuffix(source, suffix string) bool {
	return len(source) >= len(suffix) && strings.EqualFold(source[len(source)-len(suffix):], suffix)
}
