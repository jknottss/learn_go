package main

import "fmt"

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func Init() *HashTable {
	res := &HashTable{}
	for i := range res.array {
		res.array[i] = &bucket{}
	}
	return res
}

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)

}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)

}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func (b *bucket) insert(k string) {
	if !b.search(k) {
		node := &bucketNode{key: k}
		node.next = b.head
		b.head = node
	} else {
		fmt.Println(k, "already exist")
	}
}

func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	prevNode := b.head
	for prevNode.next != nil {
		if prevNode.next.key == k {
			prevNode.next = prevNode.next.next
		}
		prevNode = prevNode.next
	}
}

func main() {
	hashTable := Init()
	hashTable.Insert("Kirill")

}
