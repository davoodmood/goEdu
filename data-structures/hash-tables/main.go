package main

import "fmt"

const ArraySize = 10 // Size of the Hashtable

// HashTable structure
// will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}
// Insert will take in a key and addd it to the hash table array
func (hashtable *HashTable) insert(key string) {
	index := hash(key)
	hashtable.array[index].insert(key)
}
// Search will take in a key and return true if the bucket has that key
func (hashtable *HashTable) search(key string) bool {
	index := hash(key)
	return hashtable.array[index].search(key)
}
// Delete
func (hashtable *HashTable) delete(key string) {
	index := hash(key)
	hashtable.array[index].delete(key)
}

// bucket Structure
type bucket struct {
	head *bucketNode
}
// Insert
func (b *bucket) insert(key string) {
	if !b.search(key) {
		newNode := &bucketNode{key:key}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(key, "already exists")
	}
}
// Search will take in a key and return true if the bucket has that key
func (b *bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil{
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
// Delete
func (b *bucket) delete(key string) {
	if b.head.key == key{
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil{
		if previousNode.next.key == key {
			//delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// bucketNode structure
type bucketNode struct {
	key  string
	next *bucketNode
}

 
// hash function
func hash(key string) int{
	sum := 0
	for _,v := range key{
		sum+=int(v)
	}
	return sum % ArraySize
}

// Init
// will create a bucket in each slot of the hash table
func Init() *HashTable{
	result := &HashTable{}
	for i:= range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()

	list := []string{
		"David",
		"John",
		"Sara",
		"Nina",
		"Token",
	}
	
	for _,v := range list {
		hashTable.insert(v)
	}

	fmt.Println(hashTable.search("David"))

	// testBucket := &bucket{}
	// testBucket.insert("DAVID")
	// fmt.Println(testBucket.search("DAVID"))
	// testBucket.delete("DAVID")
	// fmt.Println(testBucket.search("DAVID"))
	
	
}