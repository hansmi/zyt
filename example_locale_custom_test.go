package zyt_test

import (
	"fmt"
	"time"

	"github.com/hansmi/zyt"
	"github.com/hansmi/zyt/pkg/zytdata"
	"golang.org/x/text/language"
)

func ExampleLocale_custom() {
	data := zytdata.LocaleData{
		Tag: language.MustParse("x-custom"),
		Months: [12]zytdata.MonthInfo{
			{
				Name:       "MyJanuary",
				Abbr:       "MyJan",
				ExtraNames: []string{"AlsoJanuary"},
			},
			// Other months omitted for brevity
		},
		Days: [7]zytdata.DayInfo{
			{
				Name: "MyMonday",
				Abbr: "MyMon",
			},
			// Other days omitted for brevity
		},
	}

	l := zyt.New(data)

	fmt.Println(l.Format("Monday, 1 January, 2006", time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC)))

	ts, err := l.ParseInLocation("January 2006", "AlsoJanuary 2001", time.UTC)
	if err != nil {
		panic(err)
	}
	fmt.Println(ts.Format(time.RFC3339))

	// Output:
	// MyMonday, 1 MyJanuary, 2001
	// 2001-01-01T00:00:00Z
}
