package zyt

import (
	"github.com/hansmi/zyt/internal/english"
	"github.com/hansmi/zyt/pkg/zytdata"
)

type localeDataMapper interface {
	mapsToEnglish() bool
	monthName(genitive bool) *mapper
	monthAbbr(genitive bool) *mapper
	dayName() *mapper
	dayAbbr() *mapper
	amPm() *mapper
}

type localeDataToEnglish zytdata.LocaleData

var _ localeDataMapper = (*localeDataToEnglish)(nil)

func (d localeDataToEnglish) mapsToEnglish() bool {
	return true
}

func (d *localeDataToEnglish) monthName(_ bool) *mapper {
	var b mapperBuilder

	for idx, mi := range d.Months {
		names := []string{mi.Name}
		if mi.GenitiveName != "" {
			names = append(names, mi.GenitiveName)
		}
		names = append(names, mi.ExtraNames...)

		b.addMultiple(names, english.MonthNames[idx])
	}

	return b.build()
}

func (d *localeDataToEnglish) monthAbbr(_ bool) *mapper {
	var b mapperBuilder

	for idx, mi := range d.Months {
		abbrs := []string{mi.Abbr}
		if mi.GenitiveAbbr != "" {
			abbrs = append(abbrs, mi.GenitiveAbbr)
		}
		abbrs = append(abbrs, mi.ExtraAbbr...)

		b.addMultiple(abbrs, english.MonthAbbrs[idx])
	}

	return b.build()
}

func (d *localeDataToEnglish) dayName() *mapper {
	var b mapperBuilder

	for idx, di := range d.Days {
		names := []string{di.Name}
		names = append(names, di.ExtraNames...)

		b.addMultiple(names, english.DayNames[idx])
	}

	return b.build()
}

func (d *localeDataToEnglish) dayAbbr() *mapper {
	var b mapperBuilder

	for idx, di := range d.Days {
		abbrs := []string{di.Abbr}
		abbrs = append(abbrs, di.ExtraAbbr...)

		b.addMultiple(abbrs, english.DayAbbrs[idx])
	}

	return b.build()
}

func (d *localeDataToEnglish) amPm() *mapper {
	var b mapperBuilder

	b.addMultiple(append([]string{d.AmPm.AM}, d.AmPm.ExtraAM...), english.AM)
	b.addMultiple(append([]string{d.AmPm.PM}, d.AmPm.ExtraPM...), english.PM)

	return b.build()
}

type localeDataFromEnglish zytdata.LocaleData

var _ localeDataMapper = (*localeDataFromEnglish)(nil)

func (d localeDataFromEnglish) mapsToEnglish() bool {
	return false
}

func (d *localeDataFromEnglish) monthName(genitive bool) *mapper {
	var b mapperBuilder

	for idx, mi := range d.Months {
		repl := mi.Name
		if genitive && mi.GenitiveName != "" {
			repl = mi.GenitiveName
		}

		b.add(english.MonthNames[idx], repl)
	}

	return b.build()
}

func (d *localeDataFromEnglish) monthAbbr(genitive bool) *mapper {
	var b mapperBuilder

	for idx, mi := range d.Months {
		repl := mi.Abbr
		if genitive && mi.GenitiveAbbr != "" {
			repl = mi.GenitiveAbbr
		}

		b.add(english.MonthAbbrs[idx], repl)
	}

	return b.build()
}

func (d *localeDataFromEnglish) dayName() *mapper {
	var b mapperBuilder

	for idx, di := range d.Days {
		b.add(english.DayNames[idx], di.Name)
	}

	return b.build()
}

func (d *localeDataFromEnglish) dayAbbr() *mapper {
	var b mapperBuilder

	for idx, di := range d.Days {
		b.add(english.DayAbbrs[idx], di.Abbr)
	}

	return b.build()
}

func (d *localeDataFromEnglish) amPm() *mapper {
	var b mapperBuilder

	b.add(english.AM, d.AmPm.AM)
	b.add(english.PM, d.AmPm.PM)

	return b.build()
}
