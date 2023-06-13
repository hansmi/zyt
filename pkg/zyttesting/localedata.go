package zyttesting

import (
	"strings"
	"testing"

	"github.com/hansmi/zyt/pkg/zytdata"
	"golang.org/x/text/language"
)

func checkAcceptable(t *testing.T, name, value string) {
	t.Helper()

	if strings.TrimSpace(value) == "" {
		t.Errorf("Attribute %q is required, got %q", name, value)
	} else if strings.TrimSpace(value) != value {
		t.Errorf("Attribute %q starts or ends with space, got %q", name, value)
	}
}

type LocaleDataTest struct{}

func (tc LocaleDataTest) Run(t *testing.T, data zytdata.LocaleData) {
	t.Helper()

	if data.Tag.IsRoot() || data.Tag == language.Und {
		t.Errorf("Language tag must be set, got %#v", data.Tag)
	}

	for _, m := range data.Months {
		t.Run(m.Name, func(t *testing.T) {
			checkAcceptable(t, "Name", m.Name)
			checkAcceptable(t, "Abbr", m.Abbr)

			for _, i := range m.ExtraNames {
				checkAcceptable(t, "ExtraNames", i)
			}

			for _, i := range m.ExtraAbbr {
				checkAcceptable(t, "ExtraAbbr", i)
			}
		})
	}

	for _, d := range data.Days {
		t.Run(d.Name, func(t *testing.T) {
			checkAcceptable(t, "Name", d.Name)
			checkAcceptable(t, "Abbr", d.Abbr)

			for _, i := range d.ExtraNames {
				checkAcceptable(t, "ExtraNames", i)
			}

			for _, i := range d.ExtraAbbr {
				checkAcceptable(t, "ExtraAbbr", i)
			}
		})
	}

	t.Run("AM/PM", func(t *testing.T) {
		ap := data.AmPm

		checkAcceptable(t, "AM", ap.AM)
		checkAcceptable(t, "PM", ap.PM)

		for _, i := range ap.ExtraAM {
			checkAcceptable(t, "ExtraAM", i)
		}

		for _, i := range ap.ExtraPM {
			checkAcceptable(t, "ExtraPM", i)
		}
	})
}
