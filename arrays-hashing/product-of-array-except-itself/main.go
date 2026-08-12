package main

import (
	"fmt"
)

// Given an integer array nums, return an array output where output[i] is the product of all the elements of nums except nums[i].
//Each product is guaranteed to fit in a 32-bit integer.

/*

	Naive solution is to go use two for loops to find the output array for each index O(n^2)

	A better solution is to use a prefix and postfix array to store the products before and after each index

	Rules:
		 - If 0 or length(nums) index of array put 1 as placeholder
		 - For every index in prefix array, it is the product of all values before it
		 - For every index in postfix array, it is the product of all values after it
		 - Output is the product: output[n] = prefix[n-1] * postfix[n+1]

	Performance:
	Complexity: O(n)
	Memory O(n)

		        _________________
i.e 	input   | 1 | 2 | 3 | 4 |
		        -----------------
			    _________________
		prefix  | 1 | 2 | 6 | 24 |
			    -----------------
		    	___________________
		postfix | 24 | 24 | 12 | 4 |
		    	-------------------
		    	 ___________________
		output   | 24 | 12 | 8 | 6 |
		    	 -------------------

	To get better performance ~memory-wise~ is to leverage a key rule from the instructions:
		The output array does not count towards memory allocation

	Given this, we can use the output as the ONLY array used besides the initial input. We can do this by using the output array as a state machine,
	first populating it with the prefix values on our fist pass over nums;
	then finding the output values on our second pass when we calculate the postfix value and then multiply by the value at nums[i] (The prefix)


	i.e.

	First pass:
					_________________
			input   | 1 | 2 | 3 | 4 |
					-----------------

					_________________
			output  | 1 | 1 | 2 | 6 |
					-----------------

	Second pass:
					_________________
			input   | 1 | 1 | 2 | 6 |
					-----------------

					___________________
			output  | 24 | 12 | 8 | 6 |
					-------------------

*/

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))

	prefix := 1
	for i, v := range nums {
		output[i] = prefix
		prefix *= v
	}
	fmt.Printf("Array after prefix: %+v\n", output)

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		output[i] *= postfix
		postfix *= nums[i]
	}
	fmt.Printf("Array after postfix: %+v\n", output)

	return output
}
