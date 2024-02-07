package pokecache

import (
	"testing"
	"time"
)



func TestCreateCache(t *testing.T) {
	interval := time.Millisecond
	cache := NewCache(interval)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	interval := time.Millisecond
	cache := NewCache(interval)

	cases := []struct {
		inputKey string
		inputVal []byte
	} {
		{
			inputKey: "key",
			inputVal: []byte("val"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "key3",
			inputVal: []byte("val3"),
		},
	}

	for _, cs := range cases {
		cache.Add(cs.inputKey, []byte(cs.inputVal))
		actual, ok := cache.Get(cs.inputKey)

		if !ok {
			t.Errorf("%s not found", cs.inputKey)
			continue
		}

		if string(actual) != string(cs.inputVal) {
			t.Errorf("%s does not match %s",
				string(actual),
				string(cs.inputVal),
			)
			continue
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"

	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	if _, ok := cache.Get(keyOne); ok {
		t.Errorf("%s should have been reaped", keyOne)
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	keyOne := "key1"

	cache.Add(keyOne, []byte("val1"))

	time.Sleep(interval / 2)

	if _, ok := cache.Get(keyOne); !ok {
		t.Errorf("%s should not have been reaped", keyOne)
	}
}
