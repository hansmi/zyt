package zyt

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestTranslator(t *testing.T) {
	for _, tc := range []struct {
		name    string
		mappers mapperSlice
		input   string
		want    string
		wantErr error
	}{
		{name: "empty"},
		{
			name: "success",
			mappers: []*mapper{
				func() *mapper {
					var b mapperBuilder
					b.add("hello", "HELLO")
					b.add("world", "unused")
					return b.build()
				}(),
				func() *mapper {
					var b mapperBuilder
					b.add("hello", "unused")
					b.add("wo", "unused")
					b.add("world", "123")
					return b.build()
				}(),
			},
			input: "prefix hello world suffix",
			want:  "prefix HELLO 123 suffix",
		},
		{
			name:    "incomplete",
			mappers: []*mapper{{}},
			input:   "hello world",
			want:    "hello world",
			wantErr: errIncompleteTranslation,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := translate(tc.input, tc.mappers)

			if diff := cmp.Diff(tc.wantErr, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("Error diff (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("translate() diff (-want +got):\n%s", diff)
			}
		})
	}
}
