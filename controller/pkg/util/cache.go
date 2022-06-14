package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"cornerstone_verifier/pkg/log"
	"cornerstone_verifier/pkg/models"

	"github.com/allegro/bigcache/v3"
)

type BigCache struct {
	cornerstoneData *bigcache.BigCache
	data            *bigcache.BigCache
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
		log.ServerError.Printf("new bigccache: %v", err)
		return nil
	}

	return &BigCache{
		cornerstoneData: bCache,
		data:            bCache,
	}
}

func dataKey(id string) string {
	return id
}

func (bc *BigCache) UpdateString(id, item string) error {
	bs, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	return bc.data.Set(dataKey(id), bs)
}

func (bc *BigCache) UpdateStruct(id string, item interface{}) error {
	bs, err := json.Marshal(&item)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	return bc.data.Set(dataKey(id), bs)
}

func (bc *BigCache) ReadString(id string) (string, error) {
	var data string
	bs, err := bc.data.Get(dataKey(id))
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return "", err
		}

		return "", fmt.Errorf("get: %w", err)
	}

	err = json.Unmarshal(bs, &data)
	if err != nil {
		return "", fmt.Errorf("unmarshal: %w", err)
	}

	return data, nil
}

func (bc *BigCache) ReadStruct(id string) (interface{}, error) {
	var data interface{}
	bs, err := bc.data.Get(dataKey(id))
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return "", err
		}

		return "", fmt.Errorf("get: %w", err)
	}

	err = json.Unmarshal(bs, &data)
	if err != nil {
		return "", fmt.Errorf("unmarshal: %w", err)
	}

	return data, nil
}

func (bc *BigCache) Delete(id string) {
	bc.data.Delete(dataKey(id))
}

func (bc *BigCache) UpdateDataCache(invitation_key string, data models.ProofRequestData) error {
	bs, err := json.Marshal(&data)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	return bc.cornerstoneData.Set(dataKey(invitation_key), bs)
}

func (bc *BigCache) ReadDataCache(invitation_key string) (models.ProofRequestData, error) {
	bs, err := bc.cornerstoneData.Get(dataKey(invitation_key))
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return models.ProofRequestData{}, err
		}

		return models.ProofRequestData{}, fmt.Errorf("get: %w", err)
	}

	var cornerstoneData models.ProofRequestData
	err = json.Unmarshal(bs, &cornerstoneData)
	if err != nil {
		return models.ProofRequestData{}, fmt.Errorf("unmarshal: %w", err)
	}

	return cornerstoneData, nil
}

func (bc *BigCache) DeleteDataCache(invitation_key string) {
	bc.cornerstoneData.Delete(dataKey(invitation_key))
}
