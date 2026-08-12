package main

import (
	"fmt"
	"testing"
)

var solution *Solution = &Solution{}

func TestMain(t *testing.T) {
	strList := [][]string{{"Hello", "World", "what's", "up"}, {"How", "Are", "We", "doing!!", "12345", "test1234!@#$%"}, {"", ""}}
	for _, v := range strList {
		s := solution.Encode(v)
		fmt.Printf("Encoded str: %s, length: %d\n", s, len(s))
		fmt.Printf("Decoded str: %+v\n", solution.Decode(s))
	}
}

//	func TestEncoding(t *testing.T) {
//		strs := []string{"hello", "human", "how", "are", "you"}
//		fmt.Println(solution.Encode(strs))
//	}
// func TestGetLength(t *testing.T) {
// 	str := "4#hello8#triangle"
// 	strR := []rune(str)
// 	length, startingIndex, _ := findLength(strR, 0)
// 	decoded := []rune{}
// 	i := startingIndex
// 	for ; i < (startingIndex+length)+1; i++ {
// 		decoded = append(decoded, strR[i])
// 	}
// 	length, startingIndex, _ = findLength(strR, i)
//
// 	i = startingIndex
// 	for ; i < (startingIndex + length)+1; i++ {
// 		decoded = append(decoded, strR[i])
// 	}
//
// 	fmt.Printf("String!: %s\n", string(decoded))
//
// }
