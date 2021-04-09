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

	fmt.Println("# Blocks in Log:", log.GetNumBlocks())

	fmt.Println("Log is valid:", log.Check())

	log.Attack(1)

	fmt.Println("After attack, Log is valid:", log.Check())
}


