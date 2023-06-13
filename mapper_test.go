package zyt

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMapper(t *testing.T) {
	var b mapperBuilder

	b.add("old", "new")
	b.add("older", "longer")

	b.add("\u00e4", "ae NFC")
	// Replace NFC variant
	b.add("a\u0308", "ae NFD")

	for _, tc := range []struct {
		name          string
		text          string
		wantRepl      string
		wantSkipCount int
	}{
		{name: "empty"},
		{
			name: "not found",
			text: "hello",
		},
		{
			name: "no match at beginning",
			text: "xold",
		},
		{
			name:          "prefix match",
			text:          "oldfoo",
			wantRepl:      "new",
			wantSkipCount: 3,
		},
		{
			name:          "longer wins",
			text:          "olderfoo",
			wantRepl:      "longer",
			wantSkipCount: 5,
		},
		{
			name:          "normalization",
			text:          "\u00e4",
			wantRepl:      "ae NFD",
			wantSkipCount: 2,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			repl, skipCount := b.build().replacePrefix(tc.text)

			if diff := cmp.Diff(tc.wantRepl, repl); diff != "" {
				t.Errorf("replacePrefix() replacement diff (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.wantSkipCount, skipCount); diff != "" {
				t.Errorf("replacePrefix() skip count diff (-want +got):\n%s", diff)
			}
		})
	}
}
