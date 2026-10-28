package memtable

import "strings"

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
	return value, true
}

func (m *MemTable) Get(key string) ([]byte, bool) {
	if len(strings.TrimSpace(key)) == 0 {
		return nil, false
	}
	val, ok := m.data[key]

	if !ok {
		return nil, false
	}
	return val, true
}

func (m *MemTable) Delete(key string) ([]byte, bool) {
	if len(strings.TrimSpace(key)) == 0 {
		return nil, false
	}
	delete(m.data, key)

	val, ok := m.Get(key)

	if !ok {
		return nil, false
	}

	return val, ok
}
