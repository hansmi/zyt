package finnish

import (
	"testing"

	"github.com/hansmi/zyt/pkg/zyttesting"
)

func TestFinnish(t *testing.T) {
	test := zyttesting.LocaleDataTest{}
	test.Run(t, Finnish)
}
