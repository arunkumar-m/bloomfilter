package main

import (
	"hash/fnv"
	"fmt"
)

type  Bloomfilter struct{
	slice []byte
	size int
	hashFunctions int
	itemsCount int
}

func getHash(s []byte) uint32 {
	h := fnv.New32()
	h.Write(s)
	return h.Sum32()
}

func getHashes(str []byte, size int) (int, int) {
	var hash1 = getHash(str)
	var hash2 = hash1 << 16;
	return int(hash1), int(hash2)
}

func newBloomfilter(size, hashFuntions int) (*Bloomfilter, error) {

	b := new(Bloomfilter)
	b.slice = make([]byte, size)
	b.size = size
	b.hashFunctions = hashFuntions
	b.itemsCount = 0

	return b, nil
}

func (b *Bloomfilter) add(value []byte) {

	hash1, hash2 := getHashes(value, b.size)

	var i int = 0
	for (i < b.hashFunctions) {
		hashValue := hash1 + i * hash2;
		b.slice[hashValue % b.size] = 1
		i = i +1
	}
	b.itemsCount++
	fmt.Println("Slice", b.slice)
}

func (b *Bloomfilter) check(val []byte) bool {

	hash1, hash2 := getHashes(val, b.size)

	var i int = 0
	for (i < b.hashFunctions) {
		hashValue := hash1 + i * hash2;
		if(b.slice[hashValue % b.size] == 0) {
			return false
		}
		i = i + 1
	}
        return true
}

func (b *Bloomfilter) clear() {
	b.slice = make([]byte, b.size)
}

func (b *Bloomfilter) numOfItems() (int) {
	return b.itemsCount
}

