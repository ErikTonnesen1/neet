package dynamicarrayimpl

import (
	"fmt"
	"testing"
)

func TestNewDynamArray(t *testing.T) {
	var da *DynamicArray = NewDynamicArray(5)
	if len(da.array) != 5 {
		t.Errorf(`Dynamic array creation resulted in incorrect length: %d`, len(da.array))
	} else if da.capacity != 5 {
		t.Errorf(`Dynamic array creation resulted in incorrect capacity: %d`, da.capacity)
	}
}

func TestGet(t *testing.T) {
	da := NewDynamicArray(5)
	da.array[3] = 13

	testVal := da.Get(3)

	if testVal != 13 {
		t.Errorf(`Test value was supposed to be 13, but was %d`, testVal)
	}
}

func TestGetAfterInit(t *testing.T) {
	da := NewDynamicArray(5)
	v := da.Get(1)
	if v != 0 {
		t.Errorf("Incorrect return value from Get (Should be 0 ): %d)", v)
	}
}

func TestPushback(t *testing.T) {
	//Given
	da := NewDynamicArray(5)
	da.array = []int{1, 2, 3, 4, 5}

	fmt.Print("Array before pushback: ", da.array)

	//When
	da.Pushback(0)

	fmt.Print("Array after pushback: ", da.array)

	//Then
	var isInitialPositionMoved bool = da.array[0] == 2
	var isLastPositionNewValue bool = da.array[len(da.array)-1] == 1

	if !isInitialPositionMoved || !isLastPositionNewValue {
		t.Errorf(`Both are to be true, but one is false: isInitialPositionMoved: %t, isLastPositionNewValue: %t`, isInitialPositionMoved, isLastPositionNewValue)
	}
}

func TestRecursivePushback(t *testing.T) {
	da := NewDynamicArray(1)
	fmt.Println("Capacity: ", da.capacity)
	da.Pushback(1)
	if da.capacity != 2 {
		t.Errorf("Double capacity didn't work, capacity turned out to be %d", da.capacity)
	}
	fmt.Println("Array: ", da.array)
}

func TestScenario2(t *testing.T) {
	//	["Array", 1, "pushback", 1, "getCapacity", "pushback", 2, "getCapacity"]
	da := NewDynamicArray(1)
	da.Pushback(1)
	fmt.Println("Capacity: ", da.GetCapacity())
	da.Pushback(2)
	fmt.Println("Capacity: ", da.GetCapacity())
}

func TestStuff(t *testing.T) {
	da := NewDynamicArray(5)
	fmt.Println("TestStuff: ")
	fmt.Println(da.array[3])

	da.Set(3, 2)
	fmt.Println(da.array[3])
	fmt.Println("Size: ", da.size)
	da.Set(3, 0)
	fmt.Println(da.array[3])
	fmt.Println("Size: ", da.size)
}
