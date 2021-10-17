package components

import (
	localLog "github.com/daqnext/LocalLog/log"
	gofastcache "github.com/daqnext/go-fast-cache"
)

func InitFastCache(localLogger_ *localLog.LocalLog) *gofastcache.LocalCache {
	return gofastcache.New(localLogger_)
}
