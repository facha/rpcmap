package main

import (
	"testing"
)

func TestMapService_PutAndGet(t *testing.T) {
	ms := NewMapService()

	// Test storing a value
	key, value := "key1", "value1"
	ms.Put(key, value)

	// Test retrieving the value
	gotValue, ok := ms.Get(key)
	if !ok {
		t.Fatalf("Get() failed, expected key '%s' to be present", key)
	}
	if gotValue != value {
		t.Errorf("Get() = %v; want %v", gotValue, value)
	}

	// Test retrieving a non-existent key
	_, ok = ms.Get("nonexistent")
	if ok {
		t.Errorf("Get() expected nonexistent key to return false")
	}
}

func TestMapService_Delete(t *testing.T) {
	ms := NewMapService()

	// Add a key-value pair to delete
	key, value := "key2", "value2"
	ms.Put(key, value)

	// Test deleting the key
	deleted := ms.Delete(key)
	if !deleted {
		t.Errorf("Delete() failed, expected key '%s' to be deleted", key)
	}

	// Ensure the key is no longer present
	_, ok := ms.Get(key)
	if ok {
		t.Errorf("Get() after Delete() expected key '%s' to be absent", key)
	}

	// Test deleting a non-existent key
	deleted = ms.Delete("nonexistent")
	if deleted {
		t.Errorf("Delete() on nonexistent key expected false, got true")
	}
}

func TestMapService_ConcurrentAccess(t *testing.T) {
	ms := NewMapService()
	key := "key3"
	value := "value3"

	// Concurrently put and get the same key
	go ms.Put(key, value)
	go func() {
		gotValue, ok := ms.Get(key)
		if ok && gotValue != value {
			t.Errorf("Concurrent Get() = %v; want %v", gotValue, value)
		}
	}()
	go ms.Delete(key)
}

