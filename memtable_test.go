package memtable

import "testing"

func TestMemTable(t *testing.T) {
	m := New()

	m.Put("name", []byte("Justine"))

	value, ok := m.Get("name")

	if !ok || string(value) != "Justine" {
		t.Fatalf("unexpected value: %q", value)
	}

	_, ok = m.Get("missing")

	if ok {
		t.Fatal("missing key should not exists")
	}
}
