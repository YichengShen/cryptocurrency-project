package main

import (
	"fmt"
	"./telog"
)

func main() {
	log := telog.Telog{}
	log.Init()

	log.AddBlock("Goofy mints 5 dollars")
	log.AddBlock("Goofy paid Alice 5 dollars")
	log.AddBlock("Alice paid Bob 5 dollars")

	fmt.Print("Log is valid: ")
	fmt.Println(log.Check())
}


