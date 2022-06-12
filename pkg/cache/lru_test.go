package cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func unpackLRU(lru *LRU) []keyValue {
	ret := make([]keyValue, 0, lru.cap)

	for e := lru.list.Front(); e != nil; e = e.Next() {
		ret = append(ret, e.Value.(keyValue))
	}

	return ret
}

func get(key string) func(*LRU) {
	return func(lru *LRU) {
		lru.Get(key)
	}
}

func set(key, value string) func(*LRU) {
	return func(lru *LRU) {
		lru.Set(key, value)
	}
}

func remove(key string) func(*LRU) {
	return func(lru *LRU) {
		lru.Remove(key)
	}
}

func TestLRU(t *testing.T) {
	tests := []struct {
		name string
		cap  int
		ops  []func(*LRU)
		want []keyValue
	}{
		{
			name: "overflow",
			cap:  2,
			ops: []func(*LRU){
				set("a", "b"),
				set("b", "b"),
				set("c", "b"),
				set("d", "b"),
				set("e", "b"),
			},
			want: []keyValue{
				{"e", "b"},
				{"d", "b"},
			},
		},
		{
			name: "reset same element",
			cap:  1,
			ops: []func(*LRU){
				set("a", "a"),
				set("a", "b"),
				set("a", "c"),
				set("a", "d"),
				set("a", "e"),
			},
			want: []keyValue{
				{"a", "e"},
			},
		},
		{
			name: "order with get calls",
			cap:  3,
			ops: []func(*LRU){
				set("a", "a"),
				set("b", "b"),
				set("c", "c"),
				get("b"),
				get("a"),
			},
			want: []keyValue{
				{"a", "a"},
				{"b", "b"},
				{"c", "c"},
			},
		},
		{
			name: "order with set calls",
			cap:  3,
			ops: []func(*LRU){
				set("a", "a"),
				set("b", "b"),
				set("c", "cc"),
				set("b", "bb"),
				set("a", "aa"),
			},
			want: []keyValue{
				{"a", "aa"},
				{"b", "bb"},
				{"c", "cc"},
			},
		},
		{
			name: "remove some elements",
			cap:  3,
			ops: []func(*LRU){
				set("a", "a"),
				set("b", "b"),
				set("c", "c"),
				remove("c"),
				remove("a"),
			},
			want: []keyValue{
				{"b", "b"},
			},
		},
		{
			name: "zero cap",
			cap:  0,
			ops: []func(*LRU){
				set("a", "a"),
				get("a"),
				remove("a"),
			},
			want: []keyValue{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lru := NewLRU(tt.cap)
			for _, op := range tt.ops {
				op(lru)
			}

			require.Equal(t, len(tt.want), lru.size)
			require.Equal(t, len(tt.want), lru.list.Len())
			require.Equal(t, len(tt.want), len(lru.cache))
			require.Equal(t, tt.want, unpackLRU(lru))
		})
	}
}
