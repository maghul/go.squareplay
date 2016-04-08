package main

import (
	"sync"
)

type syncMap struct {
	m     map[string]interface{}
	mutex sync.Mutex
}

func newSyncMap() *syncMap {
	return &syncMap{make(map[string]interface{}), sync.Mutex{}}
}

func (sm *syncMap) Get(ref string, factory func() (interface{}, error)) (interface{}, error) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	v, ok := sm.m[ref]
	if !ok {
		var err error
		v, err = factory()
		if err != nil {
			return nil, err
		}
		if v == nil {
			panic("SyncMap factory returned nil instance but no error")
		}
		sm.m[ref] = v
	}
	return v, nil
}

// Returns a snapshot of the current map.
func (sm *syncMap) snapshot() map[string]interface{} {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	rv := make(map[string]interface{})
	for r, v := range sm.m {
		rv[r] = v
	}
	return rv
}
