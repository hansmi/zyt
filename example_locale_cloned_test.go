package zyt_test

import (
	"fmt"
	"time"

	"github.com/hansmi/zyt"
)

func ExampleLocale_cloned() {
	data := zyt.English.Data().Clone()
	data.Months[2].Name = "MyMarch"

	l := zyt.New(data)

	fmt.Println(l.Format("January 2006", time.Date(2000, time.March, 1, 0, 0, 0, 0, time.UTC)))

	// Output:
	// MyMarch 2000
}
