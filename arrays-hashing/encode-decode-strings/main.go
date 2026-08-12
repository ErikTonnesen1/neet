package main

/*
* Design an algorithm to encode a list of strings to a string.
* The encoded string is then sent over the network and is decoded back to the original list of strings.
 */

/*
Need to know

  - Length of the word
  - Delimeter between length number and rest of word (to avoid) parsing a valid number in the string as a length declarator
*/

import (
	"strconv"
)

type Solution struct {
}

func (s *Solution) Encode(strs []string) string {
	var encoded string
	for _, v := range strs {
		lengthOfWord := strconv.Itoa(len(v))
		delimeter := "#"
		encoded += lengthOfWord + delimeter + v
	}
	return encoded
}

// TO DO
// Assuming the encoding contract is <string_length>#<string>
func (s *Solution) Decode(encoded string) []string {
	decoded := []string{}
	encodedRa := []rune(encoded)

	index := 0
	for index < len(encodedRa) {
		decodedStr := []rune{}
		length, base := findLength(encodedRa, index)

		if length == 0 {
			decoded = append(decoded, "")
			index += 2
		} else {
			i := base
			for ; i < (base + length); i++ {
				decodedStr = append(decodedStr, encodedRa[i])
			}
			decoded = append(decoded, string(decodedStr))
			index = i
		}
	}

	return decoded
}

func findLength(strRa []rune, startAt int) (length, startingIndex int) {
	i := startAt
	numRa := []rune{}
	for ; i < len(strRa) && strRa[i] != '#'; i++ {
		numRa = append(numRa, strRa[i])
	}

	length, err := strconv.Atoi(string(numRa))
	if err != nil || length == 0 {
		return 0, 0
	}
	return length, i + 1
}
