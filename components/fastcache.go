package components

import gofastcache "github.com/daqnext/go-fast-cache"

func InitFastCache() *gofastcache.LocalCache {
	return gofastcache.New()
}
