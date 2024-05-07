package main

import (
	piscine "01Files/sudoku/sudoku"
	"flag"
	"fmt"
	"time"
)

func main() {
	Difficulty := flag.Int("diff", 30, "How many numbers to be missing (1-64)")
	Mode := flag.Bool("hard", false, "The game won't tell you if guesses are wrong")
	flag.Parse()

	if *Difficulty < 1 || *Difficulty > 64 {
		*Difficulty = 30
	}

	var exit string
	var wrongCount int
	piscine.GameStart()
	fulldoku := piscine.Sudoku()
	sudoku, complete, miss := piscine.Setup(*Difficulty, fulldoku)
	timer := time.Now()
	gameStatus := true

	for gameStatus {
		gameStatus = piscine.GameState(sudoku, complete, miss, &exit, &wrongCount, *Mode)
	}

	if exit != "exit" {
		finalTime := time.Since(timer).Truncate(time.Second)
		fmt.Println("\nComplete!")
		if wrongCount > 0 {
			finalTime = finalTime + (time.Duration(wrongCount * 20))
			fmt.Printf("Wrong Guesses: %v x 20 seconds = %v seconds\n", wrongCount, wrongCount*20)
		}
		fmt.Println("Time Elapsed: " + time.Duration.String(finalTime))
	}

	fmt.Println("Thanks for playing!")
}
