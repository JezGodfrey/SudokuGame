package piscine

import "fmt"

func GameStart() {
	fmt.Println("-------------------------------")
	fmt.Println("|                             |")
	fmt.Println("|                             |")
	fmt.Println("|         S U D O K U         |")
	fmt.Println("|                             |")
	fmt.Println("|                             |")
	fmt.Println("-------------------------------")
	fmt.Println("\nUse the -diff flag to change difficulty from 1 - 64 (Default: 30)")
	fmt.Println("\nHow to play: Enter index and guess like so - \"XX N\" (e.g. AF 8)")
	fmt.Println("or type exit to quit")
	fmt.Println("\nLoading...")
}
