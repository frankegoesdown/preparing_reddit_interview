package main

import "container/list"

func main() {

}

type MyHashMap struct {
	mp  []*list.List
	cap int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{mp: make([]*list.List, 1000), cap: 1000}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	// Get the list. If not, create a new one
	mpK := this.getKey(key)
	if this.mp[mpK] == nil {
		this.mp[mpK] = list.New()
	}
	// Iterate the list and overwrite the old key's value.
	l := this.mp[mpK]
	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.([]int)
		if key == v[0] {
			e.Value.([]int)[1] = value
			return
		}
	}
	// If there is no old key, we create a new key-value pair
	this.mp[mpK].PushBack([]int{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	mpK := this.getKey(key)
	l := this.mp[mpK]
	if l == nil {
		return -1
	}
	// iterate the list to get the value
	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.([]int)
		if key == v[0] {
			return v[1]
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	mpK := this.getKey(key)
	l := this.mp[mpK]
	if l == nil {
		return
	}
	// check if key exists, then remove it from the list
	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.([]int)
		if key == v[0] {
			l.Remove(e)
		}
	}
}

func (this *MyHashMap) getKey(key int) int {
	return key % this.cap
}
