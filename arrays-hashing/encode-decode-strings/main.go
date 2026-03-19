package main

/*
* Design an algorithm to encode a list of strings to a string.
* The encoded string is then sent over the network and is decoded back to the original list of strings.
 */

import ()

type Solution struct {
	delimiter, key string
}

// deliminate strings by Solution.delimiter, encode strings into encoded values by key
func (s *Solution) Encode(strs []string) string {
	var encoded string
	for i, v := range strs {

	}
}

func (s *Solution) Decode(encoded string) []string {

}
