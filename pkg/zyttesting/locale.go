package zyttesting

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type Locale interface {
	ParseInLocation(layout, value string, loc *time.Location) (time.Time, error)
	Format(layout string, t time.Time) string
}

type ParseInLocationTest struct {
	Name    string
	Layout  string
	Value   string
	Loc     *time.Location
	Want    time.Time
	WantErr error
}

func (tc ParseInLocationTest) Run(t *testing.T, l Locale) {
	name := tc.Name

	if name == "" {
		name = tc.Value
	}

	t.Run(name, func(t *testing.T) {
		loc := tc.Loc

		if loc == nil {
			loc = time.UTC
		}

		got, err := l.ParseInLocation(tc.Layout, tc.Value, loc)

		if diff := cmp.Diff(tc.WantErr, err, cmpopts.EquateErrors()); diff != "" {
			t.Errorf("Error diff (-want +got):\n%s", diff)
		}

		if err == nil {
			if diff := cmp.Diff(tc.Want, got, cmpopts.EquateApproxTime(time.Millisecond)); diff != "" {
				t.Errorf("ParseInLocation(%q, %q) diff (-want +got):\n%s", tc.Layout, tc.Value, diff)
			}
		}
	})
}

type FormatTest struct {
	Name   string
	Layout string
	Value  time.Time
	Want   string
}

func (tc FormatTest) Run(t *testing.T, l Locale) {
	name := tc.Name

	if name == "" {
		name = tc.Value.String()
	}

	t.Run(name, func(t *testing.T) {
		got := l.Format(tc.Layout, tc.Value)

		if diff := cmp.Diff(tc.Want, got); diff != "" {
			t.Errorf("Format(%q, %q) diff (-want +got):\n%s", tc.Layout, tc.Value, diff)
		}

		if _, err := l.ParseInLocation(tc.Layout, got, tc.Value.Location()); err != nil {
			t.Errorf("ParseInLocation(%q, %q) failed: %v", tc.Layout, got, err)
		}
	})
}

type LocaleTest struct {
	ParseInLocation []ParseInLocationTest

	Format []FormatTest
}

func (tc LocaleTest) Run(t *testing.T, l Locale) {
	t.Helper()

	t.Run("ParseInLocation", func(t *testing.T) {
		for _, tc := range tc.ParseInLocation {
			tc.Run(t, l)
		}
	})

	t.Run("Format", func(t *testing.T) {
		for _, tc := range tc.Format {
			tc.Run(t, l)
		}
	})

	t.Run("Roundtrip", func(t *testing.T) {
		localeRoundtrip(t, l)
	})
}

func localeRoundtrip(t *testing.T, l Locale) {
	t.Helper()

	times := []time.Time{
		time.Now(),
		time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	}

	for month := time.January; month <= time.December; month++ {
		for day := 1; day <= 7; day++ {
			times = append(times, time.Date(2000, month, day,
				(int(month)+day)%24, (int(month)*day)%60, (len(times)*day)%60, 0,
				time.UTC))
		}
	}

	for _, tc := range []struct {
		layout string
		trunc  time.Duration
	}{
		{time.Layout, time.Second},
		{time.RFC822, time.Minute},
		{time.RFC850, time.Second},
		{time.RFC3339, time.Second},
		{time.RFC3339Nano, time.Nanosecond},
		{time.UnixDate, time.Second},
	} {
		for _, ts := range times {
			formatted := l.Format(tc.layout, ts)

			if parsed, err := l.ParseInLocation(tc.layout, formatted, ts.Location()); err != nil {
				t.Errorf("ParseInLocation(%q, %q) failed: %v", tc.layout, formatted, err)
			} else if diff := cmp.Diff(ts.Truncate(tc.trunc), parsed, cmpopts.EquateApproxTime(time.Millisecond)); diff != "" {
				t.Errorf("ParseInLocation(%q, %q) diff (-want +got):\n%s", tc.layout, formatted, diff)
			}
		}
	}
}
