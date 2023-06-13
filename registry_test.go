package zyt

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hansmi/zyt/internal/xtextlanguage"
	"golang.org/x/text/language"
)

func TestDefaultBest(t *testing.T) {
	for _, tc := range []struct {
		name           string
		tag            language.Tag
		want           language.Tag
		wantConfidence language.Confidence
	}{
		{
			name:           "empty",
			want:           language.English,
			wantConfidence: language.No,
		},
		{
			name:           "en-us",
			tag:            language.AmericanEnglish,
			want:           language.English,
			wantConfidence: language.Exact,
		},
		{
			name:           "en-gb",
			tag:            language.BritishEnglish,
			want:           language.English,
			wantConfidence: language.High,
		},
		{
			name:           "de",
			tag:            language.German,
			want:           language.German,
			wantConfidence: language.Exact,
		},
		{
			name:           "de-at",
			tag:            xtextlanguage.MustCompose(language.German, language.MustParseRegion("at")),
			want:           xtextlanguage.MustCompose(language.German, language.MustParseRegion("at")),
			wantConfidence: language.Exact,
		},
		{
			name:           "de-ch",
			tag:            xtextlanguage.MustCompose(language.German, language.MustParseRegion("ch")),
			want:           language.German,
			wantConfidence: language.High,
		},
		{
			name:           "fr",
			tag:            language.French,
			want:           language.English,
			wantConfidence: language.No,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, c := Best(tc.tag)

			if diff := cmp.Diff(tc.want, got.Tag(), cmp.Comparer(func(a, b language.Tag) bool {
				return a == b
			})); diff != "" {
				t.Errorf("Best() locale diff (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.wantConfidence, c); diff != "" {
				t.Errorf("Best() confidence diff (-want +got):\n%s", diff)
			}
		})
	}
}
