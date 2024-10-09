package main

import (
	"sync"
	"time"
)

type item struct {
	value               interface{}
	lastAccessInSeconds int64
	maxTTLInSeconds     int64
}

type TTLMap struct {
	m  map[string]*item
	mu sync.Mutex
}

func New(size int) *TTLMap {
	m := &TTLMap{
		m: make(map[string]*item, size),
	}

	go func(m *TTLMap) {
		for now := range time.Tick(time.Second) {
			m.mu.Lock()
			for k, v := range m.m {
				if now.Unix()-v.lastAccessInSeconds > v.maxTTLInSeconds {
					delete(m.m, k)
				}
			}
			m.mu.Unlock()
		}
	}(m)
	return m
}

func (m *TTLMap) Get(key string) (interface{}, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if it, ok := m.m[key]; ok {
		it.lastAccessInSeconds = time.Now().Unix()
		return it.value, true
	}
	return nil, false
}

func (m *TTLMap) Put(key string, value interface{}, maxTTL time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	it, ok := m.m[key]
	if !ok {
		it = &item{}
	}
	it.value = value
	it.lastAccessInSeconds = time.Now().Unix()
	it.maxTTLInSeconds = int64(maxTTL.Seconds())
	m.m[key] = it
}

func (m *TTLMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.m, key)
}

func (m *TTLMap) Len() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.m)
}
