package main

import (
	"sync"
)

type MapService struct {
	data sync.Map
}

func NewMapService() *MapService {
	return &MapService{}
}

func (ms *MapService) Put(key, value string) {
	ms.data.Store(key, value)
}

func (ms *MapService) Get(key string) (string, bool) {
	value, ok := ms.data.Load(key)
	if !ok {
		return "", false
	}
	return value.(string), true
}

func (ms *MapService) Delete(key string) bool {
	_, ok := ms.data.Load(key)
	if ok {
		ms.data.Delete(key)
	}
	return ok
}

