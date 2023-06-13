package zyt_test

import (
	"testing"
	"time"

	"github.com/hansmi/zyt"
	"github.com/hansmi/zyt/pkg/zyttesting"
)

func TestFinnish(t *testing.T) {
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
				Value:  "02 maalisk. 23 15:04 UTC",
				Want:   time.Date(2023, time.March, 2, 15, 4, 0, 0, time.UTC),
			},
			{
				Name:   "long",
				Layout: "2. January 2006",
				Value:  "12. joulukuuta 2000",
				Want:   time.Date(2000, time.December, 12, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:   "Monday",
				Layout: "Monday, 2.1.2006",
				Value:  "maanantai, 1.4.2020",
				Want:   time.Date(2020, time.April, 1, 0, 0, 0, 0, time.UTC),
			},
			{
				Name:   "Monday short",
				Layout: "Mon, 2.1.2006 PM",
				Value:  "Mo, 1.4.2020 PM",
				Want:   time.Date(2020, time.April, 1, 12, 0, 0, 0, time.UTC),
			},
			{
				Name:   "month and year",
				Layout: "January 2006",
				Value:  "elokuu 2008",
				Want:   time.Date(2008, time.August, 1, 0, 0, 0, 0, time.UTC),
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
				Value:  time.Date(2023, time.March, 2, 15, 4, 0, 0, time.UTC),
				Want:   "02 maalis 23 15:04 UTC",
			},
			{
				Name:   "long",
				Layout: "2. January 2006",
				Value:  time.Date(2000, time.August, 6, 0, 0, 0, 0, time.UTC),
				Want:   "6. elokuuta 2000",
			},
			{
				Name:   "day",
				Layout: "Mon, 2. January 2006",
				Value:  time.Date(2000, time.December, 12, 0, 0, 0, 0, time.UTC),
				Want:   "ti, 12. joulukuuta 2000",
			},
			{
				Name:   "am",
				Layout: time.Kitchen,
				Value:  time.Date(2000, time.January, 1, 1, 2, 3, 0, time.UTC),
				Want:   "1:02AM",
			},
			{
				Name:   "january",
				Layout: "2. January 2006",
				Value:  time.Date(2000, time.January, 12, 0, 0, 0, 0, time.UTC),
				Want:   "12. tammikuuta 2000",
			},
			{
				Name:   "month",
				Layout: "January 2006",
				Value:  time.Date(2001, time.March, 1, 0, 0, 0, 0, time.UTC),
				Want:   "maaliskuu 2001",
			},
		},
	}

	test.Run(t, zyt.Finnish)
}
