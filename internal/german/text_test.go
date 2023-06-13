package german

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRemoveDiacritics(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  string
	}{
		{},
		{
			input: "nfc \u00E4,\u00F6,\u00FC \u00C4,\u00D6,\u00DC",
			want:  "nfc a,o,u A,O,U",
		},
		{
			input: "nfd a\u0308,o\u0308,u\u0308 A\u0308,O\u0308,U\u0308",
			want:  "nfd a,o,u A,O,U",
		},
	} {
		t.Run(tc.input, func(t *testing.T) {
			got, err := removeDiacritics(tc.input)

			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("removeDiacritics() diff (-want +got):\n%s", diff)
			}
		})
	}
}
