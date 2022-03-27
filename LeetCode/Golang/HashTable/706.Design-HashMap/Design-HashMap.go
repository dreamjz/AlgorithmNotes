package hashtable

import (
	"container/list"
	"fmt"
	"strings"
)

const base = 769

type entry struct {
	key   int
	value int
}

type MyHashMap struct {
	data []list.List
}

func (m *MyHashMap) hash(key int) int {
	return key % base
}

func Constructor706() MyHashMap {
	return MyHashMap{data: make([]list.List, base)}
}

func (m *MyHashMap) put(key, val int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			e.Value = entry{key: key, value: val}
			return
		}
	}
	m.data[h].PushBack(entry{key: key, value: val})
}

func (m *MyHashMap) Get(key int) int {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			return et.value
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			m.data[h].Remove(e)
		}
	}
}

func (m MyHashMap) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	for _, l := range m.data {
		for e := l.Front(); e != nil; e = e.Next() {
			sb.WriteString(fmt.Sprintf("%v, ", e.Value))
		}
	}
	sb.WriteString("}")
	return sb.String()
}
