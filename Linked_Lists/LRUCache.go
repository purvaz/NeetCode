package LinkedLists

import "fmt"

type CacheNode struct {
	Key   int
	Value int
	Left  *CacheNode
	Right *CacheNode
}

type LRUCache struct {
	Capacity int
	CacheMap map[int]*CacheNode
	LRU      *CacheNode
	MRU      *CacheNode
}

var length int

func Constructor(capacity int) LRUCache {
	LRUObj := &LRUCache{
		Capacity: capacity,
		CacheMap: make(map[int]*CacheNode),
		LRU:      &CacheNode{Key: 0, Value: 0},
		MRU:      &CacheNode{Key: 0, Value: 0},
	}
	LRUObj.LRU.Right = LRUObj.MRU
	LRUObj.MRU.Left = LRUObj.LRU
	return *LRUObj
}

func (this *LRUCache) Get(key int) int {
	fmt.Println("Getting values for : ", key)
	if _, exists := this.CacheMap[key]; exists {
		if this.MRU != this.CacheMap[key] {
			this.Remove(this.CacheMap[key])
			this.Insert(this.CacheMap[key])
		}
		DisplayDoublyLinkedList(this.LRU)
		return this.CacheMap[key].Value
	}
	DisplayDoublyLinkedList(this.LRU)
	return -1
}

func (this *LRUCache) Remove(node *CacheNode) {
	left, right := node.Left, node.Right
	left.Right, right.Left = right, left
}

func (this *LRUCache) Insert(node *CacheNode) {
	if this.LRU == node {
		this.LRU = this.LRU.Right
	}
	node.Left = this.MRU
	node.Right = nil
	this.MRU.Right = node
	this.MRU = node
}

func (this *LRUCache) Put(key int, value int) {
	fmt.Println("Putting values : ", key, value)

	// fmt.Println("--------------------", this.CacheMap)
	// if it already exists?
	if _, exists := this.CacheMap[key]; exists {
		// fmt.Println("This is the existing value : ", this.CacheMap[key].Value)
		// update the value in Node
		this.CacheMap[key].Value = value
		// fmt.Println("This is the updated value : ", this.CacheMap[key].Value)

		// shift the node to end
		if this.MRU != this.CacheMap[key] {
			this.Remove(this.CacheMap[key])
			this.Insert(this.CacheMap[key])
		}
	} else {
		// create the node
		obj := &CacheNode{
			Key:   key,
			Value: value,
		}
		// check length < capacity
		if length < this.Capacity {
			// Add the node in the list
			this.Insert(obj)

			// Add node in cache
			this.CacheMap[key] = obj

			// Check if first node, update LRU
			if length == 0 {
				this.LRU = obj
			}
			// increment length
			length++
		} else {
			fmt.Println("More than capacity")
			if this.Capacity > 1 {
				LRU := this.LRU
				// remove LRU from linkedlist
				this.Remove(LRU)
				delete(this.CacheMap, LRU.Key)

				// update LRU to LRU.Next
				this.LRU = LRU.Right

				// Add a node at the end
				this.Insert(obj)

				// update cache
				this.CacheMap[key] = obj
			} else {
				delete(this.CacheMap, this.LRU.Key)
				this.LRU = obj
				this.MRU = obj
				this.CacheMap[key] = obj

			}
		}
	}
	DisplayDoublyLinkedList(this.LRU)
}

func DisplayDoublyLinkedList(head *CacheNode) {

	fmt.Print("Displaying LL: ")
	current := head
	for current != nil {
		fmt.Print(current.Key, " : ", current.Value, " -> ")
		current = current.Right
	}
	fmt.Println("")
}
