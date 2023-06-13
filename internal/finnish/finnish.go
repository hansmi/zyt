package finnish

import (
	"github.com/hansmi/zyt/pkg/zytdata"
	"golang.org/x/text/language"
)

// A note to native Finnish speakers: The author implemented Finnish as an
// example of a language with genitive cases. Spelling and language syntax
// errors would not be surprising.

var Finnish = zytdata.LocaleData{
	Tag: language.Finnish,
	Months: [12]zytdata.MonthInfo{
		{
			Name:         "tammikuu",
			GenitiveName: "tammikuuta",
			Abbr:         "tammik.",
			GenitiveAbbr: "tammi",
			ExtraAbbr:    []string{"T"},
		},
		{
			Name:         "helmikuu",
			GenitiveName: "helmikuuta",
			Abbr:         "helmik.",
			GenitiveAbbr: "helmi",
			ExtraAbbr:    []string{"H"},
		},
		{
			Name:         "maaliskuu",
			GenitiveName: "maaliskuuta",
			Abbr:         "maalisk.",
			GenitiveAbbr: "maalis",
			ExtraAbbr:    []string{"M"},
		},
		{
			Name:         "huhtikuu",
			GenitiveName: "huhtikuuta",
			Abbr:         "huhtik.",
			GenitiveAbbr: "huhti",
			ExtraAbbr:    []string{"H"},
		},
		{
			Name:         "toukokuu",
			GenitiveName: "toukokuuta",
			Abbr:         "toukok.",
			GenitiveAbbr: "touko",
			ExtraAbbr:    []string{"T"},
		},
		{
			Name:         "kesäkuu",
			GenitiveName: "kesäkuuta",
			Abbr:         "kesäk.",
			GenitiveAbbr: "kesä",
			ExtraNames:   []string{"kesakuu", "kesakuuta"},
			ExtraAbbr:    []string{"K", "kesak.", "kesa"},
		},
		{
			Name:         "heinäkuu",
			GenitiveName: "heinäkuuta",
			Abbr:         "heinäk.",
			GenitiveAbbr: "heinä",
			ExtraNames:   []string{"heinakuu", "heinakuuta"},
			ExtraAbbr:    []string{"H", "heinak.", "heina"},
		},
		{
			Name:         "elokuu",
			GenitiveName: "elokuuta",
			Abbr:         "elok.",
			GenitiveAbbr: "elo",
			ExtraAbbr:    []string{"E"},
		},
		{
			Name:         "syyskuu",
			GenitiveName: "syyskuuta",
			Abbr:         "syysk.",
			GenitiveAbbr: "syys",
			ExtraAbbr:    []string{"S"},
		},
		{
			Name:         "lokakuu",
			GenitiveName: "lokakuuta",
			Abbr:         "lokak.",
			GenitiveAbbr: "loka",
			ExtraAbbr:    []string{"L"},
		},
		{
			Name:         "marraskuu",
			GenitiveName: "marraskuuta",
			Abbr:         "marrask.",
			GenitiveAbbr: "marras",
			ExtraAbbr:    []string{"M"},
		},
		{
			Name:         "joulukuu",
			GenitiveName: "joulukuuta",
			Abbr:         "jouluk.",
			GenitiveAbbr: "joulu",
			ExtraAbbr:    []string{"J"},
		},
	},
	Days: [7]zytdata.DayInfo{
		{Name: "maanantai", Abbr: "mo"},
		{Name: "tiistai", Abbr: "ti"},
		{Name: "keskiviikko", Abbr: "ke"},
		{Name: "torstai", Abbr: "to"},
		{Name: "perjantai", Abbr: "pe"},
		{Name: "lauantai", Abbr: "la"},
		{Name: "sunnuntai", Abbr: "su"},
	},
	AmPm: zytdata.AmPmInfo{
		AM: "AM",
		PM: "PM",
	},
}
