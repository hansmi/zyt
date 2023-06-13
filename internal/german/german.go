package german

import (
	"sort"
	"strings"

	"github.com/hansmi/zyt/internal/xtextlanguage"
	"github.com/hansmi/zyt/pkg/zytdata"
	"golang.org/x/exp/slices"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/norm"
)

var germanUmlautReplacer = strings.NewReplacer(
	"\u00E4", "ae",
	"\u00F6", "oe",
	"\u00FC", "ue",
	"\u00C4", "AE",
	"\u00D4", "OE",
	"\u00DC", "UE",
)

func prepareMonth(m zytdata.MonthInfo) zytdata.MonthInfo {
	m.ExtraAbbr = append(m.ExtraAbbr, m.Abbr)

	for _, i := range []struct {
		values *[]string
		addDot bool
	}{
		{&m.ExtraNames, false},
		{&m.ExtraAbbr, true},
	} {
		values := *i.values

		for _, v := range values[:] {
			if modified, err := removeDiacritics(v); err != nil {
				panic(err)
			} else if modified != v {
				values = append(values, modified)
			}

			if modified := germanUmlautReplacer.Replace(norm.NFC.String(v)); modified != v {
				values = append(values, modified)
			}
		}

		if i.addDot {
			for _, v := range values[:] {
				values = append(values, v+".")
			}
		}

		sort.Strings(values)

		*i.values = slices.Compact(values)
	}

	m.Abbr += "."

	return m
}

func prepareDay(d zytdata.DayInfo) zytdata.DayInfo {
	d.ExtraAbbr = append(d.ExtraAbbr, d.Abbr)

	for _, v := range d.ExtraAbbr[:] {
		d.ExtraAbbr = append(d.ExtraAbbr, v+".")
	}

	sort.Strings(d.ExtraAbbr)

	d.ExtraAbbr = slices.Compact(d.ExtraAbbr)

	return d
}

var German = zytdata.LocaleData{
	Tag: language.German,
	Months: func() [12]zytdata.MonthInfo {
		months := [12]zytdata.MonthInfo{
			{
				Name: "Januar", Abbr: "Jan",
				ExtraNames: []string{"Jänner"},
				ExtraAbbr:  []string{"Jän"},
			},
			{
				Name: "Februar", Abbr: "Feb",
				ExtraNames: []string{"Feber"},
				ExtraAbbr:  []string{"Febr"},
			},
			{
				Name: "März", Abbr: "Mär",
				ExtraAbbr: []string{"Mrz"},
			},
			{Name: "April", Abbr: "Apr"},
			{Name: "Mai", Abbr: "Mai"},
			{Name: "Juni", Abbr: "Jun"},
			{Name: "Juli", Abbr: "Jul"},
			{Name: "August", Abbr: "Aug"},
			{
				Name: "September", Abbr: "Sep",
				ExtraAbbr: []string{"Sept"},
			},
			{Name: "Oktober", Abbr: "Okt"},
			{Name: "November", Abbr: "Nov"},
			{Name: "Dezember", Abbr: "Dez"},
		}

		for idx, m := range months {
			months[idx] = prepareMonth(m)
		}

		return months
	}(),
	Days: func() [7]zytdata.DayInfo {
		days := [7]zytdata.DayInfo{
			{
				Name:      "Montag",
				Abbr:      "Mo",
				ExtraAbbr: []string{"Mon"},
			},
			{
				Name:      "Dienstag",
				Abbr:      "Di",
				ExtraAbbr: []string{"Die"},
			},
			{
				Name:      "Mittwoch",
				Abbr:      "Mi",
				ExtraAbbr: []string{"Mit", "Mittw"},
			},
			{
				Name:      "Donnerstag",
				Abbr:      "Do",
				ExtraAbbr: []string{"Don"},
			},
			{
				Name:      "Freitag",
				Abbr:      "Fr",
				ExtraAbbr: []string{"Fre"},
			},
			{
				Name:      "Samstag",
				Abbr:      "Sa",
				ExtraAbbr: []string{"Sam"},
			},
			{
				Name:      "Sonntag",
				Abbr:      "So",
				ExtraAbbr: []string{"Son"},
			},
		}

		for idx, d := range days {
			days[idx] = prepareDay(d)
		}

		return days
	}(),
	AmPm: zytdata.AmPmInfo{
		AM:      "vorm.",
		PM:      "nachm.",
		ExtraAM: []string{"am", "a.m.", "vorm"},
		ExtraPM: []string{"pm", "p.m.", "nachm"},
	},
}

var AustrianGerman = func() zytdata.LocaleData {
	d := German.Clone()
	d.Tag = xtextlanguage.MustCompose(d.Tag, language.MustParseRegion("AT"))

	january := &d.Months[0]
	january.Name = "Jänner"
	january.Abbr = "Jän."

	return d
}()
