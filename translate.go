package zyt

import (
	"errors"
	"fmt"
	"strings"
)

var errIncompleteTranslation = errors.New("translation incomplete")

// Apply mappers to text without overlapping in the order they're given. An
// error is returned if not all mappers find a piece of text to replace.
func translate(text string, mappers mapperSlice) (string, error) {
	var buf strings.Builder

	buf.Grow(len(text))

	textPos := 0
	mapperPos := 0

	for textPos < len(text) {
		if mapperPos >= len(mappers) {
			break
		}

		m := mappers[mapperPos]

		if repl, skipCount := m.replacePrefix(text[textPos:]); skipCount > 0 {
			buf.WriteString(repl)
			textPos += skipCount
			mapperPos++
			continue
		}

		buf.WriteByte(text[textPos])
		textPos++
	}

	if textPos < len(text) {
		buf.WriteString(text[textPos:])
	}

	var err error

	if mapperPos < len(mappers) {
		err = fmt.Errorf("%w: used only %d of %d mappers for %q with result %q",
			errIncompleteTranslation,
			mapperPos, len(mappers), text, buf.String())
	}

	return buf.String(), err
}
