package main

import (
	"fmt"
	"./telog"
)

func main() {
	head := telog.AddBlock("0", telog.CreateEmptyHashPointer())
	head1 := telog.AddBlock("1", head)
	head2 := telog.AddBlock("2", head1)
	head3 := telog.AddBlock("3", head2)
	fmt.Println(telog.Check(head3))
}


