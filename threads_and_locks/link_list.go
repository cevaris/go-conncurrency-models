package main

import (
	"sync"
)


type Node struct {
	Value int64
	Prev *Node
	Next *Node
	Mutex *sync.Mutex
}

func NewEmptyNode() *Node {
	return &Node{
		Mutex: &sync.Mutex{},
	}
}

func NewNode(value int64, prev *Node, next *Node) *Node {
	return &Node{
		Value: value,
		Prev: prev,
		Next: next,
		Mutex: &sync.Mutex{},
	}
}


type ConcurrentSortedList struct {
	Head *Node
	Tail *Node
}

func NewConcurrentSortedList() *ConcurrentSortedList {
	ls := &ConcurrentSortedList{
		Head: NewEmptyNode(),
		Tail: NewEmptyNode(),
	}
	ls.Head.Next = ls.Tail
	ls.Tail.Prev = ls.Head
	return ls
}

func (ls *ConcurrentSortedList) Insert(value int64) {
	current := ls.Head
	current.Mutex.Lock()

	next := current.Next	
	for {
		next.Mutex.Lock()
		
		if next == ls.Tail || next.Value < value {
			node := NewNode(value, current, next)
			next.Prev = node
			current.Next = node

			current.Mutex.Unlock()
			next.Mutex.Unlock()
			return
		}
		current.Mutex.Unlock()
		current = next
		next = current.Next
	}
	next.Mutex.Unlock()
}


func (ls *ConcurrentSortedList) Size() int64 {
	var count int64 = 0
	current := ls.Head
	for current.Next != ls.Tail {
		mutex := current.Mutex
		mutex.Lock()
		count += 1
		current = current.Next
		mutex.Unlock()
	}
	return count
}
