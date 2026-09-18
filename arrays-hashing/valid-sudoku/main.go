package main

import ()

type Square struct {
	x, y int
}

func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]struct{})
	cols := make(map[int]struct{})
	sqrs := make(map[coordinate]struct{})

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if _, sqrOk := sqrs[coordinate{r / 3, c / 3}]; sqrOk {
				return false
			}
			if _, rowOk := rows[r]; rowOk {
				return false
			}
			if _, colOk := cols[c]; colOk {
				return false
			}
		}
	}

}
