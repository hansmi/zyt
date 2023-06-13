package zytdata

import (
	"golang.org/x/exp/slices"
	"golang.org/x/text/language"
)

type MonthInfo struct {
	// Long month name in the grammatical form used when the month is a part of
	// a complete date.
	Name string

	// Abbreviated month name in the grammatical form used when the month is
	// a part of a complete date.
	Abbr string

	// Long month name in the grammatical form used when the month is named by
	// itself.
	GenitiveName string

	// Abbreviated month name in the grammatical form used when the month is
	// named by itself.
	GenitiveAbbr string

	// Additional month names recognized during parsing.
	ExtraNames []string

	// Additional abbreviated month names recognized during parsing.
	ExtraAbbr []string
}

// Clone returns a deep copy which can be modified without affecting the
// original.
func (mi MonthInfo) Clone() MonthInfo {
	mi.ExtraNames = slices.Clone(mi.ExtraNames)
	mi.ExtraAbbr = slices.Clone(mi.ExtraAbbr)
	return mi
}

type DayInfo struct {
	// Long-named day of the week.
	Name string

	// Abbreviated day of the week.
	Abbr string

	// Additional day names recognized during parsing.
	ExtraNames []string

	// Additional abbreviated day names recognized during parsing.
	ExtraAbbr []string
}

// Clone returns a deep copy which can be modified without affecting the
// original.
func (di DayInfo) Clone() DayInfo {
	di.ExtraNames = slices.Clone(di.ExtraNames)
	di.ExtraAbbr = slices.Clone(di.ExtraAbbr)
	return di
}

type AmPmInfo struct {
	// Ante meridiem.
	AM string

	// Post meridiem.
	PM string

	// Additional tokens recognized for ante meridiem.
	ExtraAM []string

	// Additional tokens recognized for post meridiem.
	ExtraPM []string
}

// Clone returns a deep copy which can be modified without affecting the
// original.
func (api AmPmInfo) Clone() AmPmInfo {
	api.ExtraAM = slices.Clone(api.ExtraAM)
	api.ExtraPM = slices.Clone(api.ExtraPM)
	return api
}

// https://sourceware.org/git?p=glibc.git;a=blob;f=locale/LocaleData.h;hb=HEAD
type LocaleData struct {
	Tag    language.Tag
	Months [12]MonthInfo
	Days   [7]DayInfo
	AmPm   AmPmInfo
}

// Clone returns a deep copy which can be modified without affecting the
// original.
func (d LocaleData) Clone() LocaleData {
	for idx, mi := range d.Months {
		d.Months[idx] = mi.Clone()
	}

	for idx, di := range d.Days {
		d.Days[idx] = di.Clone()
	}

	d.AmPm = d.AmPm.Clone()

	return d
}
