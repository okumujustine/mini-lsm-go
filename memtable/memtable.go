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

	if len(strings.TrimSpace(key)) == 0 {
		return nil, false
	}
	cloneValue := bytes.Clone(value)
	m.data[key] = cloneValue

	return cloneValue, true
}

func (m *MemTable) Get(key string) ([]byte, bool) {
	if len(strings.TrimSpace(key)) == 0 {
		return nil, false
	}
	val, ok := m.data[key]
	if !ok {
		return nil, false
	}

	if len(val) == 0 {
		return nil, false
	}
	return bytes.Clone(val), true
}

func (m *MemTable) Delete(key string) ([]byte, bool) {
	if len(strings.TrimSpace(key)) == 0 {
		return nil, false
	}

	val, ok := m.Put(key, make([]byte, 0))

	return val, ok
}
