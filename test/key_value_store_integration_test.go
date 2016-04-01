package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/myteksi/kvstore"
)

var result kvstore.KVStore
var valueSave int

func TestReadWriteDeleteKVStore(t *testing.T) {
	kv := kvstore.NewKVStore()
	kv.Set("key", 100)
	kv.Set("key2", "Value")

	value, _ := kv.Get("key")
	assert.Equal(t, 100, value)

	kv.Delete("key")
	value, ok := kv.Get("key")
	assert.Nil(t, value)
	assert.False(t, ok)
}

func BenchmarkKVStoreWrite(b *testing.B) {
	kv := kvstore.NewKVStore()
	for i := 0; i < b.N; i++ {
		kv.Set(string(i), i)
	}
	result = kv
}

func BenchmarkKVStoreRead(b *testing.B) {
	kv := kvstore.NewKVStore()
	for i := 0; i < b.N; i++ {
		kv.Set(string(i), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		value, _ := kv.Get(string(i))
		valueSave = value.(int)
	}
}

func BenchmarkKVStoreDelete(b *testing.B) {
	kv := kvstore.NewKVStore()
	for i := 0; i < b.N; i++ {
		kv.Set(string(i), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		kv.Delete(string(i))
	}
	result = kv
}

func benchmarkSingleRoutine(kv kvstore.KVStore, b *testing.B, done chan<- bool) {
	for i := 0; i < b.N; i++ {
		key := string(i % 100)
		kv.Set(key, i)
		value, _ := kv.Get(key)
		valueSave, _ = value.(int)
		kv.Delete(key)
	}
	done <- true
}

func benchmarkConcurrent(count int, b *testing.B) {
	done := make(chan bool)
	kv := kvstore.NewKVStore()
	for i := 0; i < count; i++ {
		go benchmarkSingleRoutine(kv, b, done)
	}

	for i := 0; i < count; i++ {
		<-done
	}
}

func BenchmarkConcurrent2(b *testing.B) {
	benchmarkConcurrent(2, b)
}

func BenchmarkConcurrent10(b *testing.B) {
	benchmarkConcurrent(10, b)
}
