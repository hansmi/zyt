package english

import (
	"github.com/hansmi/zyt/pkg/zytdata"
	"golang.org/x/text/language"
)

var MonthNames = [12]string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

var MonthAbbrs = [12]string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var DayNames = [7]string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

var DayAbbrs = [7]string{
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
	"Sun",
}

const AM = "AM"
const PM = "PM"

var English = zytdata.LocaleData{
	Tag: language.English,
	Months: func() [12]zytdata.MonthInfo {
		var months [12]zytdata.MonthInfo

		for idx := range months {
			m := &months[idx]
			m.Name = MonthNames[idx]
			m.Abbr = MonthAbbrs[idx]
			m.ExtraAbbr = append(m.ExtraAbbr, m.Abbr+".")
		}

		months[8].ExtraAbbr = []string{"Sept", "Sept."}

		return months
	}(),
	Days: func() [7]zytdata.DayInfo {
		var days [7]zytdata.DayInfo

		for idx := range days {
			d := &days[idx]
			d.Name = DayNames[idx]
			d.Abbr = DayAbbrs[idx]
			d.ExtraAbbr = append(d.ExtraAbbr, d.Abbr[:2])
		}

		days[1].ExtraAbbr = []string{"Tues"}
		days[2].ExtraAbbr = []string{"Weds"}
		days[3].ExtraAbbr = []string{"Thur", "Thurs"}

		for idx := range days {
			d := &days[idx]
			for _, abbr := range d.ExtraAbbr {
				d.ExtraAbbr = append(d.ExtraAbbr, abbr+".")
			}
			d.ExtraAbbr = append(d.ExtraAbbr, d.Abbr+".")
		}

		return days
	}(),
	AmPm: zytdata.AmPmInfo{
		AM:      AM,
		PM:      PM,
		ExtraAM: []string{"a.m."},
		ExtraPM: []string{"p.m."},
	},
}
