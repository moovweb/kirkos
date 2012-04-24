package kirkos

import "time"
import "errors"

type KirkosValue struct {
	key string
	value Value
	stamp int64
}

type KirkosCache struct {
	size int
	reclaim int
	cache map[string]*KirkosValue
}

func NewKirkosCache(size int, reclaim int) Cache {
	c := KirkosCache{size: size, reclaim: reclaim, cache: map[string]*KirkosValue{}}
	return Cache(c)
}

func (c KirkosCache) deleteItem(item *KirkosValue) {
	delete(c.cache, item.key)
	defer func() {
		recover()
	}()
	ok := item.value.(Freeable)
	if ok != nil {
		ok.Free()
	}
}

func (c KirkosCache) reclaimStorage() {
	for i := 0; i < c.reclaim; i++ {
		oldest := &KirkosValue{stamp: time.Now().UnixNano(), value: nil}
		for _, item := range c.cache {
			if item.stamp < oldest.stamp {
				oldest = item
			}
		}
		c.deleteItem(oldest)
	}
}

func (c KirkosCache) Get(key string) (Value, error) {
	if c.cache[key] != nil {
		return c.cache[key].value, nil
	}
	return nil, errors.New("Invalid key")
}

func (c KirkosCache) Set(key string, value Value) {
	if len(c.cache) >= c.size {
		c.reclaimStorage()
	}
	c.cache[key] = &KirkosValue{key: key, value: value, stamp: time.Now().UnixNano()}
}

func (c KirkosCache) Free() {
	for _, item := range c.cache {
		c.deleteItem(item)
	}
}

