package utils

import (
	"encoding/json"
	"errors"
	"time"

	"iamza-verifier/pkg/log"

	"github.com/allegro/bigcache/v3"
)

type BigCache struct {
	stringData *bigcache.BigCache
}

func NewBigCache() *BigCache {
	bCache, err := bigcache.NewBigCache(bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,

		// time after which entry can be evicted
		LifeWindow: 8760 * time.Hour, // 1 year

		// Interval between removing expired entries (clean up).
		// If set to <= 0 then no action is performed.
		// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
		CleanWindow: 5 * time.Minute,

		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,

		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,

		// prints information about additional memory allocation
		Verbose: false,

		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 256,

		// callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		OnRemove: nil,

		// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
		// for the new entry, or because delete was called. A constant representing the reason will be passed through.
		// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
		// Ignored if OnRemove is specified.
		OnRemoveWithReason: nil,
	})
	if err != nil {
		log.ServerError.Printf("Failed on NewBigCache(): %s", err.Error())
		return nil
	}

	return &BigCache{
		stringData: bCache,
	}
}

func (bc *BigCache) String(key, entry string) error {
	bs, err := json.Marshal(entry)
	if err != nil {
		log.ServerError.Printf("Failed to marshal string cache: %s", err.Error())
	}

	return bc.stringData.Set(key, bs)
}

func (bc *BigCache) ReadString(key string) (string, error) {
	bs, err := bc.stringData.Get(key)
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			log.ServerInfo.Println("Cache entry not found")
		}

		// log.ServerError.Printf("Failed to read string cache: %s", err.Error())
	}

	var entryData string
	err = json.Unmarshal(bs, &entryData)
	if err != nil {
		// log.ServerError.Printf("Failed to unmarshal string cache: %s", err.Error())
	}

	return entryData, nil
}

func (bc *BigCache) DeleteString(key string) {
	bc.stringData.Delete(key)
}
