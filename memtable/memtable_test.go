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

func TestDelete(t *testing.T) {
	m := New()

	m.Put("user:1", []byte("Alice"))
	m.Put("user:2", []byte("Bob"))

	m.Delete("user:1")

	_, ok := m.Get("user:1")
	if ok {
		t.Fatal("deleted key should not be found")
	}

	value, ok := m.Get("user:2")
	if !ok || string(value) != "Bob" {
		t.Fatal("other keys should remain intact")
	}

	// A newer write should replace the tombstone.
	m.Put("user:1", []byte("Charlie"))

	value, ok = m.Get("user:1")
	if !ok || string(value) != "Charlie" {
		t.Fatal("key should exist after reinsertion")
	}
}

func TestPutDoesNotExposeMemory(t *testing.T) {
	m := New()

	original := []byte("Alice")

	returned, ok := m.Put("name", original)
	if !ok {
		t.Fatal("Put failed")
	}

	// Mutate the original slice.
	original[0] = 'M'

	// Mutate the returned slice.
	returned[0] = 'B'

	stored, ok := m.Get("name")
	if !ok || string(stored) != "Alice" {
		t.Fatalf("stored value was modified: %q", stored)
	}
}
