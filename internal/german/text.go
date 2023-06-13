package german

import (
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var diacriticsTransformer = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

// Remove non-spacing marks, in particular the two dots placed above German
// umlauts.
func removeDiacritics(text string) (string, error) {
	result, _, err := transform.String(diacriticsTransformer, text)

	return result, err
}
