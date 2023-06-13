package zyttesting

import (
	"testing"
	"time"
)

type stdTime struct{}

func (stdTime) ParseInLocation(layout, value string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(layout, value, loc)
}

func (stdTime) Format(layout string, t time.Time) string {
	return t.Format(layout)
}

func TestStdTime(t *testing.T) {
	test := LocaleTest{}
	test.Run(t, &stdTime{})
}
