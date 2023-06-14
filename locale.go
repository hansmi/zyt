package zyt

import (
	"fmt"
	"time"
	"unicode"

	"github.com/hansmi/zyt/pkg/zytdata"
	"golang.org/x/text/language"
	"golang.org/x/text/unicode/norm"
)

const (
	monthNameToken = "January"
	monthAbbrToken = "Jan"

	dayNameToken = "Monday"
	dayAbbrToken = "Mon"

	pmToken = "PM"
)

// Break up a time layout string at the boundaries between digits, letters and
// punctuation.
//
// TODO: Consider using the rune before and after a token for more precise
// matching, but only if it's not a number or a layout token itself (e.g.
// recognize spaces and punctuation only).
func walkLayout(text string, handler func(string)) {
	var prevCategories [3]bool
	var begin int

	textRunes := []rune(text)

	for idx, r := range textRunes {
		runeCategories := [3]bool{
			unicode.IsDigit(r),
			unicode.IsLetter(r),
			unicode.IsPunct(r),
		}

		if idx > 0 && prevCategories != runeCategories {
			handler(string(textRunes[begin:idx]))
			begin = idx
		}

		prevCategories = runeCategories
	}

	if begin <= len(textRunes) {
		handler(string(textRunes[begin:]))
	}
}

func resolveToken(lm localeDataMapper, token string, genitive bool) *mapper {
	switch token {
	case monthNameToken:
		return lm.monthName(genitive)
	case monthAbbrToken:
		return lm.monthAbbr(genitive)
	case dayNameToken:
		return lm.dayName()
	case dayAbbrToken:
		return lm.dayAbbr()
	case pmToken:
		return lm.amPm()
	}

	return nil
}

func isNumericDayToken(token string) bool {
	switch token {
	case "2", "_2", "02":
		return true
	}

	return false
}

func parseLayout(lm localeDataMapper, layout string) mapperSlice {
	var result mapperSlice
	var delayedMonth map[int]string
	var numericDay bool

	walkLayout(layout, func(token string) {
		if token == monthNameToken || token == monthAbbrToken {
			if delayedMonth == nil {
				// Most layouts have only one month-related token.
				delayedMonth = make(map[int]string, 1)
			}

			delayedMonth[len(result)] = token
			result = append(result, nil)
		} else if m := resolveToken(lm, token, false); m != nil {
			result = append(result, m)
		} else if !numericDay {
			numericDay = isNumericDayToken(token)
		}
	})

	for idx, token := range delayedMonth {
		result[idx] = resolveToken(lm, token, numericDay)
	}

	return result
}

type Locale struct {
	tag  language.Tag
	data *zytdata.LocaleData

	toEnglish   localeDataMapper
	fromEnglish localeDataMapper
}

// New instantiates a new locale.
func New(d zytdata.LocaleData) *Locale {
	return &Locale{
		tag:         d.Tag,
		data:        &d,
		toEnglish:   (*localeDataToEnglish)(&d),
		fromEnglish: (*localeDataFromEnglish)(&d),
	}
}

func (l *Locale) String() string {
	return l.tag.String()
}

func (l *Locale) Tag() language.Tag {
	return l.tag
}

func (l *Locale) Data() zytdata.LocaleData {
	return *l.data
}

func (l *Locale) parseLayoutCached(lm localeDataMapper, layout string) mapperSlice {
	key := cacheKey{
		tag:       l.tag.String(),
		toEnglish: lm.mapsToEnglish(),
		layout:    layout,
	}

	cache := getCache()

	if t, ok := cache.Get(key); ok {
		return t
	}

	result := parseLayout(lm, layout)

	cache.Add(key, result)

	return result
}

// ParseInLocation parses a formatted string and returns the time value it
// represents. Wraps [time.ParseInLocation].
func (l *Locale) ParseInLocation(layout, value string, loc *time.Location) (time.Time, error) {
	layout = norm.NFC.String(layout)
	value = norm.NFC.String(value)

	var wrapError func(error) error

	translated, err := translate(value, l.parseLayoutCached(l.toEnglish, layout))

	if err != nil {
		return time.Time{}, err
	}

	if translated != value {
		wrapError = func(err error) error {
			return fmt.Errorf("after translating %q to %q for locale %q: %w",
				value, translated, l.String(), err)
		}
	}

	parsed, err := time.ParseInLocation(layout, translated, loc)

	if !(err == nil || wrapError == nil) {
		err = wrapError(err)
	}

	return parsed, err
}

// Format returns a textual representation of the time value formatted
// according to the layout defined by the argument. Wraps
// [time.Time.Format].
//
// When the layout contains a numeric form of the day and the month name, the
// genitive form of the month name is used if the locale has one.
func (l *Locale) Format(layout string, t time.Time) string {
	layout = norm.NFC.String(layout)

	formatted, err := translate(t.Format(layout), l.parseLayoutCached(l.fromEnglish, layout))

	if err != nil {
		panic(err)
	}

	return formatted
}
