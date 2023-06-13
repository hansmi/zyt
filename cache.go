package zyt

import (
	"sync"

	lru "github.com/hashicorp/golang-lru/v2"
)

const DefaultCacheSize int = 16

type cacheKey struct {
	tag       string
	toEnglish bool
	layout    string
}

var layoutCacheOnce sync.Once
var layoutCache *lru.Cache[cacheKey, mapperSlice]

func getCache() *lru.Cache[cacheKey, mapperSlice] {
	layoutCacheOnce.Do(func() {
		cache, err := lru.New[cacheKey, mapperSlice](DefaultCacheSize)
		if err != nil {
			panic(err)
		}

		layoutCache = cache
	})

	return layoutCache
}

// SetCacheSize modifies the size of the internal cache used to store
// previously parsed time layouts. The cache is shared across all locales and
// has a default size of [DefaultCacheSize].
func SetCacheSize(size int) {
	if size < 0 {
		size = 0
	}

	getCache().Resize(size)
}
