package piscine

import (
	"math/rand/v2"
)

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

// To be used against a slice of 3 slices
func sqrChecker(rs [][]int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if rs[0][i] == rs[1][j] {
				return true
			}

			if rs[0][i] == rs[2][j] {
				return true
			}

			if rs[1][i] == rs[2][j] {
				return true
			}
		}
	}

	for i := 3; i < 6; i++ {
		for j := 3; j < 6; j++ {
			if rs[0][i] == rs[1][j] {
				return true
			}

			if rs[0][i] == rs[2][j] {
				return true
			}

			if rs[1][i] == rs[2][j] {
				return true
			}
		}
	}

	for i := 6; i < 9; i++ {
		for j := 6; j < 9; j++ {
			if rs[0][i] == rs[1][j] {
				return true
			}

			if rs[0][i] == rs[2][j] {
				return true
			}

			if rs[1][i] == rs[2][j] {
				return true
			}
		}
	}

	return false
}

func dupeChecker(rs [][]int, r []int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < len(rs); j++ {
			if rs[j][i] == r[i] {
				return true
			}
		}
	}

	return false
}

// func FinalRow(sdk [][]int) []int {
// 	var lastRow []int
// 	for i := 0; i < 9; i++ {
// 		lastNum := 45
// 		for j := 0; j < 8; j++ {
// 			lastNum = lastNum - sdk[j][i]
// 		}
// 		lastRow = append(lastRow, lastNum)
// 	}
// 	return lastRow
// }

func genRows() [][]int {
	var su [][]int
	var do []int = genRow()
	var ku []int
	var temp [][]int
	var dupe bool
	var sqrCheck = false
	su = append(su, do)
	for i := 0; i < 8; i++ {
		if sqrCheck {
			i = i - 3
			sqrCheck = false
		}

		ku = genRow()
		dupe = dupeChecker(su, ku)

		if dupe {
			i--
			continue
		}

		temp = nil
		switch len(su) {
		case 3:
			temp = su
			dupe = sqrChecker(temp)
		case 6:
			temp = su[3:]
			dupe = sqrChecker(temp)
		case 9:
			temp = su[6:]
			dupe = sqrChecker(temp)
		}

		if dupe {
			sqrCheck = true
			su = su[:len(su)-2]
			continue
		}

		if !dupe {
			su = append(su, ku)
		}
	}

	// su = append(su, FinalRow(su))
	return su
}

// Generate a row of numbers from 1-9 randomly
func genRow() []int {
	var n []int
	lastNum := 45
	for i := 0; i < 8; i++ {
		dupe := false
		num := randRange(1, 10)
		for _, v := range n {
			if num == v {
				dupe = true
				break
			}
		}
		if dupe {
			i--
		} else {
			n = append(n, num)
		}
	}

	for _, v := range n {
		lastNum = lastNum - v
	}

	n = append(n, lastNum)

	return n
}

func Sudoku() []int {
	var numbers []int
	rows := genRows()

	for _, row := range rows {
		numbers = append(numbers, row...)
	}

	return numbers
}
