package zyt_test

import (
	"testing"
	"time"

	"github.com/hansmi/zyt"
	"github.com/hansmi/zyt/pkg/zyttesting"
)

func TestEnglish(t *testing.T) {
	test := zyttesting.LocaleTest{
		ParseInLocation: []zyttesting.ParseInLocationTest{
			{
				Name:   "rfc3339",
				Layout: time.RFC3339,
				Value:  "2023-02-03T04:05:06+01:00",
				Want:   time.Date(2023, time.February, 3, 3, 5, 6, 0, time.UTC),
			},
			{
				Name:   "rfc822",
				Layout: time.RFC822,
				Value:  "02 Mar 23 15:04 UTC",
				Want:   time.Date(2023, time.March, 2, 15, 4, 0, 0, time.UTC),
			},
			{
				Name:   "long",
				Layout: "2. January 2006",
				Value:  "12. December 2000",
				Want:   time.Date(2000, time.December, 12, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:   "Monday short",
				Layout: "Mon, 2.1.2006 PM",
				Value:  "Mo, 1.4.2020 PM",
				Want:   time.Date(2020, time.April, 1, 12, 0, 0, 0, time.UTC),
			},
		},
		Format: []zyttesting.FormatTest{
			{
				Name:   "rfc3339",
				Layout: time.RFC3339,
				Value:  time.Date(2023, time.February, 3, 3, 5, 6, 0, time.UTC),
				Want:   "2023-02-03T03:05:06Z",
			},
			{
				Name:   "rfc822",
				Layout: time.RFC822,
				Value:  time.Date(2023, time.January, 2, 15, 4, 0, 0, time.UTC),
				Want:   "02 Jan 23 15:04 UTC",
			},
			{
				Name:   "long",
				Layout: "2 January 2006",
				Value:  time.Date(2000, time.March, 12, 0, 0, 0, 0, time.UTC),
				Want:   "12 March 2000",
			},
		},
	}

	test.Run(t, zyt.English)
}
