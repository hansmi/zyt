package zyt

import (
	"slices"
	"sync"

	"golang.org/x/text/language"
)

type registeredLocale struct {
	priority int
	locale   *Locale
}

type Registry struct {
	mu    sync.Mutex
	items []*registeredLocale
	tags  []language.Tag
}

// Register a locale. The locale with the lowest priority number becomes the
// default.
func (r *Registry) Register(priority int, l *Locale) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.items = append(r.items, &registeredLocale{
		priority: priority,
		locale:   l,
	})
	r.tags = nil
}

// Best selects the most appropriate locale for the requested language tag. If
// none is found the default is returned.
func (r *Registry) Best(tag language.Tag) (*Locale, language.Confidence) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.items) == 0 {
		return nil, language.No
	}

	if r.tags == nil {
		slices.SortStableFunc(r.items, func(a, b *registeredLocale) int {
			if a.priority < b.priority {
				return -1
			} else if a.priority > b.priority {
				return +1
			}

			return 0
		})

		r.tags = make([]language.Tag, len(r.items))

		for idx, i := range r.items {
			r.tags[idx] = i.locale.Tag()
		}
	}

	m := language.NewMatcher(r.tags)

	_, index, confidence := m.Match(tag)

	return r.items[index].locale, confidence
}

// The global default registry. Built-in languages are automatically
// registered.
var DefaultRegistry = &Registry{}

var Register = DefaultRegistry.Register
var Best = DefaultRegistry.Best

func init() {
	DefaultRegistry.Register(0, English)

	DefaultRegistry.Register(50, Finnish)
	DefaultRegistry.Register(50, German)

	DefaultRegistry.Register(100, AustrianGerman)
}
