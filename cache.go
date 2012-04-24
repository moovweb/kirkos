package kirkos

type Freeable interface {
	Free()
}

type Value interface {
}

type Cache interface {
	Get(string) (Value, error)
	Set(string, Value)
	Free()
}

func NewCache(size int, reclaim int) Cache {
	return NewKirkosCache(size, reclaim)
}

func NewStableCache(size int) Cache {
	return NewKirkosCache(size, 1)
}

func NewFastCache(size int) Cache {
	return NewKirkosCache(size, size / 10)
}

