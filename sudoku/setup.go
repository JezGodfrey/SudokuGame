package piscine

import (
	"fmt"
	"strconv"
)

func Setup(n int, su []int) ([]string, []string, []int) {
	var nudoku []string
	var strdoku []string
	var miss []int
	var start string
	for i := 0; i < n; i++ {
		dupeCheck := false
		genNum := randRange(0, 81)

		for _, v := range miss {
			if genNum == v {
				dupeCheck = true
			}
		}

		if dupeCheck {
			i--
			continue
		}

		miss = append(miss, genNum)
	}

	for i, v := range su {
		missCheck := false
		for _, m := range miss {
			if i == m {
				missCheck = true
				nudoku = append(nudoku, " ")
			}
		}

		if !missCheck {
			nudoku = append(nudoku, strconv.Itoa(v))
		}
	}

	for _, c := range su {
		strdoku = append(strdoku, strconv.Itoa(c))
	}

	fmt.Printf("\033[1A\033[K")
	fmt.Println("Loaded!\nPress Enter to start!")
	fmt.Scanf("%v\n", &start)

	return nudoku, strdoku, miss
}
