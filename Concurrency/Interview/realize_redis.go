package main

import (
	"errors"
	"sync"
)

type (
	SyncRedis struct {
		data map[InputData]OutputData
		mu   sync.RWMutex
	}

	InputData  string
	OutputData string

	ResultData struct {
		data OutputData
		err  error
	}
)

var (
	ErrNotFound = errors.New("not found")
)

func (rd *SyncRedis) Init() *SyncRedis {
	return &SyncRedis{
		data: make(map[InputData]OutputData),
		mu:   sync.RWMutex{},
	}
}

func (rd *SyncRedis) Put(key InputData, data OutputData) {
	rd.mu.Lock()
	defer rd.mu.Unlock()

	rd.data[key] = data
}

func (rd *SyncRedis) Get(key InputData) ResultData {
	rd.mu.RLock()
	defer rd.mu.RUnlock()
	if val, ok := rd.data[key]; ok {
		return ResultData{val, nil}
	}
	return ResultData{err: ErrNotFound}
}

func main() {

}
