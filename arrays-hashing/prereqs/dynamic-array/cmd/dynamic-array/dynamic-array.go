package dynamicarray

import (
// "fmt"
)

type DynamicArray struct {
	size  int
	array []int
}

func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		capacity,
		make([]int, capacity*2),
	}
}

func (da *DynamicArray) Get(i int) int {
	return da.array[i]
}

func (da *DynamicArray) Set(i int, n int) {
	da.size++
	da.array[i] = n
}

func (da *DynamicArray) Pushback(n int) {
	if da.size+1 >= cap(da.array) {
		da.resize()
	}
	da.array[da.size] = da.array[n]
	da.array[n] = 0
	da.size++
}

func (da *DynamicArray) Popback() int {
	lastValue := da.array[da.size-1]
	da.array[da.size-1] = 0
	da.size--
	return lastValue
}

func (da *DynamicArray) resize() {
	newSlice := make([]int, da.size, cap(da.array)*2)
	copy(newSlice, da.array)
	da.array = newSlice
}

func (da *DynamicArray) GetSize() int {
	return da.size
}

func (da *DynamicArray) GetCapacity() int {
	return cap(da.array)
}
