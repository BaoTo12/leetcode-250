package main

type Node struct {
	key  int
	val  int
	next *Node
}

type MyHashMap struct {
	buckets []*Node
	size    int
}

func Constructor() *MyHashMap {
	return &MyHashMap{
		buckets: make([]*Node, 1000),
		size:    1000,
	}
}

// put
func (m *MyHashMap) Put(key int, val int) {
	bucket := m.key(key)
	head := m.buckets[bucket]
	if head == nil {
		head = &Node{
			key:  key,
			val:  val,
			next: nil,
		}
		return
	}
	cur := head
	for {
		if cur.key == key {
			cur.val = val
			return
		}
		if cur.next == nil {
			cur.next = &Node{key: key, val: val}
			return
		}
		cur = cur.next
	}
}

// get
func (m *MyHashMap) Get(key int) int {
	bucket := m.key(key)
	cur := m.buckets[bucket]

	for cur != nil {
		if cur.key == key {
			return cur.val
		}
		cur = cur.next
	}
	return -1
}

// remove
func (m *MyHashMap) Remove(key int) {
	idx := m.key(key)
	cur := m.buckets[idx]
	if cur == nil {
		return
	}
	// If head is the key, remove head
	if cur.key == key {
		m.buckets[idx] = cur.next
		return
	}
	prev := cur
	cur = cur.next
	for cur != nil {
		if cur.key == key {
			prev.next = cur.next
			return
		}
		prev, cur = cur, cur.next
	}
}

// key
func (m *MyHashMap) key(key int) int {
	return key % len(m.buckets)
}
