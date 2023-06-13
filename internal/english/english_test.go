package english

import (
	"testing"

	"github.com/hansmi/zyt/pkg/zyttesting"
)

func TestEnglish(t *testing.T) {
	test := zyttesting.LocaleDataTest{}
	test.Run(t, English)
}
