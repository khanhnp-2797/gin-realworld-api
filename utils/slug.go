package utils

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// MakeSlug - Tạo slug từ string
func MakeSlug(s string) string {
	// Convert to lowercase
	s = strings.ToLower(s)

	// Remove accents
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ = transform.String(t, s)

	// Replace spaces and special characters with hyphens
	reg := regexp.MustCompile("[^a-z0-9]+")
	s = reg.ReplaceAllString(s, "-")

	// Remove leading/trailing hyphens
	s = strings.Trim(s, "-")

	return s
}
