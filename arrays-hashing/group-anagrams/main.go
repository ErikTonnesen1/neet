package main

import "sort"

//Read and learn this: https://siongui.github.io/2017/05/07/go-sort-string-slice-of-rune/
//Could do better, answer only beats a certain % of answers

func badSolution(strs []string) [][]string {

	if len(strs) <= 1 {
		return [][]string{strs}
	}

	//k: alphabetically sorted string, v: array of strings
	anagramMap := make(map[string][]string)
	for _, str := range strs {
		sortedString := sortString(str)
		if _, found := anagramMap[sortedString]; found == true {
			anagramMap[sortedString] = append(anagramMap[sortedString], str)
		} else {
			anagramMap[sortedString] = []string{str}
		}
	}

	var groupedAnagrams [][]string
	for _, anagramArray := range anagramMap {
		groupedAnagrams = append(groupedAnagrams, anagramArray)
	}
	return groupedAnagrams
}

func sortString(str string) string {
	var runeString []rune
	for _, v := range str {
		runeString = append(runeString, v)
	}
	//Learn this
	sort.Slice(runeString, func(i, j int) bool {
		return runeString[i] < runeString[j]
	})
	return string(runeString)
}

//Better solution
// - inline the sort, build the map first, return the map
// - finding the anagram (the anagram check) can be implicit
// Steps:
// - sort string and add to map
// - return values of map

func groupAnagrams(str []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, v := range str {
		r := []rune(v)
		sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
		key := string(r)
		anagramMap[key] = append(anagramMap[key], v)
	}

	var result [][]string
	for _, v := range anagramMap {
		result = append(result, v)
	}
	return result
}
