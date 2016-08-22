package main

import (
	"errors"
	"sync"
)

type syncMap struct {
	m     map[string]interface{}
	mutex sync.Mutex
	sn    *map[string]interface{}
}

func newSyncMap() *syncMap {
	return &syncMap{make(map[string]interface{}), sync.Mutex{}, nil}
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
		sm.sn = nil
	}
	return v, nil
}
func (sm *syncMap) Remove(ref string) (interface{}, error) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	v, ok := sm.m[ref]
	if ok {
		delete(sm.m, ref)
		sm.sn = nil
		return v, nil
	}
	return nil, errors.New("No such entry")
}

// Returns a snapshot of the current map.
func (sm *syncMap) snapshot() map[string]interface{} {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	if sm.sn != nil {
		return *sm.sn
	}
	rv := make(map[string]interface{})
	for r, v := range sm.m {
		rv[r] = v
	}
	sm.sn = &rv
	return rv
}
