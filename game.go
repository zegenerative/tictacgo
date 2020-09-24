package main

import (
	f "fmt"
	b "bufio"
	"os"
	s "strings"
)

var turn = 0

func game() {
	reader := b.NewReader(os.Stdin)
	f.Println("What are the coordinates of your choice?")
	f.Println("(Please type two integers)")
	
	f.Print("-> ")
  choice, _ := reader.ReadString('\n')
	choice = s.Replace(choice, "\n", "", -1)
	var x int
	var y int
	for i, _ := range choice {
		if i == 0 { 
			x = i
		} else {
			y = i
		}	
	}
	updateBoard([]int{x, y}, "b")
	turn += 1
	f.Println(turn)
}