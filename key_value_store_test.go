package kvstore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateKVStore(t *testing.T) {
	kv := NewKVStore()
	assert.Implements(t, (*KVStore)(nil), kv)
}

func TestSetAndGetKVStore(t *testing.T) {
	kv := NewKVStore()
	kv.Set("key1", 100)
	value, ok := kv.Get("key1")
	assert.Equal(t, 100, value)
	assert.True(t, ok)
}

func TestGetNotExistKeyShouldReturnFalse(t *testing.T) {
	kv := NewKVStore()
	value, ok := kv.Get("key1")
	assert.Nil(t, value)
	assert.False(t, ok)
}

func TestSetWithSameKeyShouldReplace(t *testing.T) {
	kv := NewKVStore()
	kv.Set("key1", "value1")
	value, ok := kv.Get("key1")
	assert.Equal(t, "value1", value)
	assert.True(t, ok)

	kv.Set("key1", "value2")
	value, ok = kv.Get("key1")
	assert.Equal(t, "value2", value)
	assert.True(t, ok)
}

func TestDeleteExistingKey(t *testing.T) {
	kv := NewKVStore()
	kv.Set("key", 100)
	value, ok := kv.Get("key")
	assert.Equal(t, 100, value)
	assert.True(t, ok)

	kv.Delete("key")
	value, ok = kv.Get("key")
	assert.Nil(t, value)
	assert.False(t, ok)
}

func TestDeleteNonexistingKeyShouldNotFail(t *testing.T) {
	kv := NewKVStore()
	kv.Set("key", 100)
	kv.Delete("non key")

	value, ok := kv.Get("key")
	assert.Equal(t, 100, value)
	assert.True(t, ok)
}
