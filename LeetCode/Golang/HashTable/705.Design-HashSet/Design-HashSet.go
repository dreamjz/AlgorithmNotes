package hashtbale

import (
	"container/list"
	"fmt"
	"log"
	"strings"
)

const base = 769

type MyHashSet struct {
	data []list.List
}

func Constructor705() MyHashSet {
	return MyHashSet{data: make([]list.List, base)}
}

func (hs *MyHashSet) hash(key int) int {
	return key % base
}

func (hs *MyHashSet) Contains(key int) bool {
	h := hs.hash(key)
	for e := hs.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}

func (hs *MyHashSet) Add(key int) {
	if !hs.Contains(key) {
		h := hs.hash(key)
		hs.data[h].PushBack(key)
	}
}

func (hs *MyHashSet) Remove(key int) {
	h := hs.hash(key)
	for e := hs.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			hs.data[h].Remove(e)
		}
	}
}

func (hs MyHashSet) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	for _, l := range hs.data {
		for e := l.Front(); e != nil; e = e.Next() {
			log.Printf("E: val: %v, Type: %T\n", e.Value, e.Value)
			sb.WriteString(fmt.Sprintf("%v, ", e.Value))
		}
	}
	sb.WriteString("}")
	return sb.String()
}
