package main

import (
	"fmt"
)

func main() {
	//Capacity , Number of hashes
	bloomFilter, _ := newBloomfilter(50, 4)

	bloomFilter.add([]byte("Hello"))
	bloomFilter.add([]byte("World"))

	fmt.Println(bloomFilter.check([]byte("Hello")))
	fmt.Println(bloomFilter.check([]byte("World1")))
	fmt.Println(bloomFilter.numOfItems())

	bloomFilter.clear()
	fmt.Println(bloomFilter.check([]byte("Hello")))

}

