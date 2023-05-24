package types

import (
	"math/rand"
	"sync"
)

type UniqueObjects[K comparable, V any] struct {
	storage     map[K]V
	mx          *sync.RWMutex
	createKeyFn func() K
}

func NewUniqueObjects[K comparable, V any](createKeyFn func() K) UniqueObjects[K, V] {
	return UniqueObjects[K, V]{
		storage:     map[K]V{},
		mx:          &sync.RWMutex{},
		createKeyFn: createKeyFn,
	}
}

func NewUniqueObjectsWithInt64Key[V any]() UniqueObjects[int64, V] {
	return UniqueObjects[int64, V]{
		storage:     map[int64]V{},
		mx:          &sync.RWMutex{},
		createKeyFn: rand.Int63,
	}
}

func (u UniqueObjects[K, V]) Add(value V) K {
	u.mx.Lock()
	defer u.mx.Unlock()

	for {
		key := u.createKeyFn()
		_, exists := u.storage[key]
		if !exists {
			u.storage[key] = value
			return key
		}
	}
}

func (u UniqueObjects[K, _]) RemoveWithKey(key K) {
	u.mx.Lock()
	delete(u.storage, key)
	u.mx.Unlock()
}

func (u UniqueObjects[_, _]) RemoveAll() {
	u.mx.Lock()
	for key := range u.storage {
		delete(u.storage, key)
	}
	u.mx.Unlock()
}

func (u UniqueObjects[_, V]) RemoveAllWithFn(beforeDeleteFn func(V)) {
	u.mx.Lock()
	for key, value := range u.storage {
		beforeDeleteFn(value)
		delete(u.storage, key)
	}
	u.mx.Unlock()
}

func (u UniqueObjects[K, V]) GetWithKey(id K) (V, bool) {
	u.mx.RLock()
	result, ok := u.storage[id]
	u.mx.RUnlock()
	return result, ok
}

func (u UniqueObjects[K, V]) AllObjectsEnumerated(ch chan V) {
	u.mx.RLock()
	for _, value := range u.storage {
		ch <- value
	}
	u.mx.RUnlock()
}

func (u UniqueObjects[K, V]) AllObjects() []V {
	u.mx.RLock()
	result := make([]V, len(u.storage))
	i := 0
	for _, value := range u.storage {
		result[i] = value
		i++
	}
	u.mx.RUnlock()
	return result
}

func (u UniqueObjects[K, V]) AllKeysEnumerated(ch chan K) {
	u.mx.RLock()
	for key := range u.storage {
		ch <- key
	}
	u.mx.RUnlock()
}

func (u UniqueObjects[K, V]) AllKeys() []K {
	u.mx.RLock()
	result := make([]K, len(u.storage))
	i := 0
	for key := range u.storage {
		result[i] = key
		i++
	}
	u.mx.RUnlock()
	return result
}
