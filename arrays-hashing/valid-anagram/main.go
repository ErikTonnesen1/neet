package main

func isAnagram(s, t string) (isAnagram bool) {

	if len(s) != len(t) {
		return false
	}

	sWordMap := make(map[rune]int)
	for _, v := range s {
		sWordMap[v]++
	}

	tWordMap := make(map[rune]int)
	for _, v := range t {
		tWordMap[v]++
	}

	for k, _ := range sWordMap {
		if sWordMap[k] != tWordMap[k] {
			return false
		}
	}

	return true
}
