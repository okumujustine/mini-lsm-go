package memtable

import (
	"bytes"
	"strings"
)

type MemTable struct {
	data map[string][]byte
}

func New() *MemTable {
	return &MemTable{
		data: make(map[string][]byte),
	}
}

func (m *MemTable) Put(key string, value []byte) ([]byte, bool) {

	if len(strings.TrimSpace(key)) == 0 || len(value) == 0 {
		return nil, false
	}
	m.data[key] = value
	return bytes.Clone(value), true
}

func (m *MemTable) Get(key string) ([]byte, bool) {
	if len(strings.TrimSpace(key)) == 0 {
		return nil, false
	}
	val, ok := m.data[key]

	if !ok {
		return nil, false
	}
	return bytes.Clone(val), true
}

func (m *MemTable) Delete(key string) ([]byte, bool) {
	if len(strings.TrimSpace(key)) == 0 {
		return nil, false
	}

	val, ok := m.Get(key)

	if !ok {
		return nil, false
	}

	delete(m.data, key)

	return val, ok
}
