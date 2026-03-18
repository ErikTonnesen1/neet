package dynamicarrayimpl

import "slices"

type DynamicArray struct {
	array          []int
	size, capacity int
}

func NewDynamicArray(capacity int) *DynamicArray {
	if capacity < 0 {
		return nil
	}

	return &DynamicArray{
		make([]int, capacity),
		0,
		capacity,
	}
}

// Assume that index i is valid.
func (da *DynamicArray) Get(i int) int {
	return da.array[i]
}

// Assume that index i is valid.
func (da *DynamicArray) Set(i int, n int) {
	if da.array[i] == 0 {
		da.size++
	}
	da.array[i] = n
}

func (da *DynamicArray) Pushback(n int) {
	if da.size+1 >= da.capacity {
		da.resize()
		da.Pushback(n)
	}
	for i := n; i < len(da.array)-1; i++ {
		temp := da.array[i]
		da.array[i] = da.array[i+1]
		da.array[i+1] = temp
	}
}

func (da *DynamicArray) Popback() int {
	length := len(da.array)
	pop := da.array[length-1]
	_ = slices.Delete(da.array, length-1, length)
	return pop
}

func (da *DynamicArray) resize() {
	newCapacity := da.capacity * 2
	resizedSlice := make([]int, da.size, newCapacity)
	copy(resizedSlice, da.array)

	da.array = resizedSlice
	da.capacity = newCapacity
}

func (da *DynamicArray) GetSize() int {
	return da.size
}

func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}
