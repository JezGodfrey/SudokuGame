package piscine

import (
	"fmt"
	"strconv"
	"strings"
)

func GameState(s []string, c []string, m []int, e *string, wrong *int, mode bool) bool {
	var key string
	var index int
	var guess int
	var Cyan = "\033[36m"
	var Red = "\033[31m"
	var Green = "\033[32m"
	var Def = "\033[0m"
	mp := make(map[string]int)
	win := false
	f := false
	colorCheck := false
	keys := []string{"AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI",
		"BA", "BB", "BC", "BD", "BE", "BF", "BG", "BH", "BI",
		"CA", "CB", "CC", "CD", "CE", "CF", "CG", "CH", "CI",
		"DA", "DB", "DC", "DD", "DE", "DF", "DG", "DH", "DI",
		"EA", "EB", "EC", "ED", "EE", "EF", "EG", "EH", "EI",
		"FA", "FB", "FC", "FD", "FE", "FF", "FG", "FH", "FI",
		"GA", "GB", "GC", "GD", "GE", "GF", "GG", "GH", "GI",
		"HA", "HB", "HC", "HD", "HE", "HF", "HG", "HH", "HI",
		"IA", "IB", "IC", "ID", "IE", "IF", "IG", "IH", "II"}

	fmt.Println("\n    A  B  C   D  E  F   G  H  I")
	fmt.Println("   -----------------------------")
	fmt.Printf("A |")
	lets := [9]string{"B", "C", "D", "E", "F", "G", "H", "I"}
	for i := 0; i < len(s); i++ {
		for _, v := range m {
			if i == v {
				colorCheck = true
			}
		}

		if colorCheck {
			if mode {
				fmt.Printf(" %v%v%v ", Cyan, s[i], Def)
			} else {
				if s[i] == c[i] {
					fmt.Printf(" %v%v%v ", Green, s[i], Def)
				} else {
					fmt.Printf(" %v%v%v ", Red, s[i], Def)
					*wrong++
				}
			}
			colorCheck = false
		} else {
			fmt.Printf(" %v ", s[i])
		}

		if (i+1)%3 == 0 {
			fmt.Printf("|")
		}

		if (i+1)%27 == 0 {
			fmt.Printf("\n   -----------------------------")
		}

		if (i+1)%9 == 0 && i != 80 {
			fmt.Printf("\n%v |", lets[i/9])
		}
	}

	for i := 0; i < 81; i++ {
		if s[i] != c[i] {
			win = true
		}
	}

	if !win {
		return win
	}

	for i := 0; i < len(keys); i++ {
		mp[keys[i]] = i
	}

	fmt.Printf("\n\"Index Guess\": ")
	fmt.Scanf("%v %v\n", &key, &guess)
	key = strings.ToUpper(key)

	if strings.ToLower(key) == "exit" {
		*e = strings.ToLower(key)
		return false
	}

	if _, ok := mp[key]; !ok {
		fmt.Println("\nInvalid input - Index needs to be AA-II")
		return true
	}

	index = mp[key]

	if guess < 1 || guess > 9 {
		fmt.Println("\nInvalid input - Guess needs to be 1-9")
		return true
	}

	for _, spc := range m {
		if index == spc {
			f = true
		}
	}

	if f {
		s[index] = strconv.Itoa(guess)
	} else {
		fmt.Println("\nInvalid input - Index of fixed number")
	}

	return true
}
