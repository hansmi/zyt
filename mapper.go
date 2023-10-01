package zyt

import (
	"strings"

	"golang.org/x/exp/slices"
	"golang.org/x/text/unicode/norm"
)

func makeMapperKey(text string) string {
	return strings.ToLower(norm.NFC.String(text))
}

type mapperBuilder struct {
	entries map[string]string
}

func (b *mapperBuilder) addMultiple(values []string, new string) {
	for _, v := range values {
		b.add(v, new)
	}
}

func (b *mapperBuilder) add(old string, new string) {
	if b.entries == nil {
		b.entries = map[string]string{}
	}

	b.entries[makeMapperKey(old)] = new
}

func (b mapperBuilder) build() *mapper {
	m := &mapper{
		entries: make([]mapperEntry, 0, len(b.entries)),
	}

	for old, new := range b.entries {
		if m.minLength == 0 || len(old) < m.minLength {
			m.minLength = len(old)
		}

		m.entries = append(m.entries, mapperEntry{old, new})
	}

	// Sort longest entries first
	slices.SortFunc(m.entries, func(a, b mapperEntry) int {
		if len(a.old) > len(b.old) {
			return -1
		} else if len(a.old) < len(b.old) {
			return +1
		}

		return strings.Compare(a.old, b.old)
	})

	return m
}

type mapperSlice []*mapper

type mapperEntry struct {
	old, new string
}

type mapper struct {
	minLength int
	entries   []mapperEntry
}

// Find a suitable replacement for the beginning of text. The first entry whose
// old value is a case-insensitive prefix of text is returned along with the
// number of bytes to skip in text.
func (m mapper) replacePrefix(text string) (string, int) {
	if len(text) < m.minLength {
		return "", 0
	}

	for _, entry := range m.entries {
		if len(entry.old) > len(text) {
			continue
		}

		if prefix := text[0:len(entry.old)]; strings.EqualFold(prefix, entry.old) {
			return entry.new, len(prefix)
		}
	}

	return "", 0
}
