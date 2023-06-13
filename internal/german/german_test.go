package german

import (
	"testing"

	"github.com/hansmi/zyt/pkg/zyttesting"
)

func TestGerman(t *testing.T) {
	test := zyttesting.LocaleDataTest{}
	test.Run(t, German)
}

func TestAustrianGerman(t *testing.T) {
	test := zyttesting.LocaleDataTest{}
	test.Run(t, AustrianGerman)
}
