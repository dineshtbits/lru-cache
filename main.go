package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	Data   int
	KeyPtr *list.Element
}

type LRUCache struct {
	Queue    *list.List
	Items    map[int]*Node
	Capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Queue: list.New(), Items: make(map[int]*Node), Capacity: capacity}
}

func (l *LRUCache) Get(key int) int {
	if item, ok := l.Items[key]; ok {
		l.Queue.MoveToFront(item.KeyPtr)
		return item.Data
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if item, ok := l.Items[key]; !ok {
		if l.Capacity == len(l.Items) {
			back := l.Queue.Back()
			l.Queue.Remove(back)
			delete(l.Items, back.Value.(int))
		}
		l.Items[key] = &Node{Data: value, KeyPtr: l.Queue.PushFront(key)}
	} else {
		item.Data = value
		l.Items[key] = item
		l.Queue.MoveToFront(item.KeyPtr)
	}

}

func main() {
	// Test case 1
	// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	//	[null, null, null, 1, null, -1, null, -1, 3, 4]
	fmt.Println("Test case 1")
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))

	// Test case 2
	// 	["LRUCache","put","put","put","put","get","get"]
	// [[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
	// [null,null,null,null,null,-1,3]
	fmt.Println("Test case 2")
	obj = Constructor(2)
	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(2, 3)
	obj.Put(4, 1)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
}
