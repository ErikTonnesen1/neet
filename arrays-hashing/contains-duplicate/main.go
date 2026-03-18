package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 5}

	var hasDuplicate bool = hasDuplicate(nums)

	fmt.Printf("Has Duplicate is %t", hasDuplicate)
}

// Given a list of ints, return true if contains a duplicate

/* Ideas:
 1. Add all to map O(N)
	- If exists return true
	- If not, add to map

 2. Use a Hashmap --> A Class that offers O(1) performance because of hashing
 	- Same logic as above
	- Really would still be O(N) overall perf, as worst case would be 0->N would have to lookup the Map
*/

/*
	Notes about Maps

	- Structure of maps = map[keyType]valueType
	- Maps are reference types, like that of pointers or slices. A nil map behaves like a empty map when reading. But attempts to write to a nil map will cause run-time panic. Do not do that.
	- Create a new map using the make function
*/

func hasDuplicate(nums []int) bool {
	var hashMap = make(map[int]int)
	hasDuplicate := false

	for _, v := range nums {
		// initialize the first value as _ to not retrieve the value, but only check if it existst
		_, found := hashMap[v]
		if found {
			hasDuplicate = true
		} else {
			hashMap[v] = v
		}
	}
	return hasDuplicate
}
