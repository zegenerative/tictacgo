package main

import (
	f "fmt"
	b "bufio"
	"os"
	s "strings"
	"strconv"
)

var turn = 1
var x int
var y int

func gameStatus() {
	checkRows(board)
	if turn == 10 {
		f.Println("It's a draw!")
		gameOver = true
	}
}

func choose() {
	f.Println("turn:", turn)
	reader := b.NewReader(os.Stdin)
	text := "Player " + player + ", what are the x and y coordinates of your choice?"
	f.Printf(text)
	f.Println("(Please type two integers (0 - 2), separated by a comma)")
	
	validAnswer := false

	for validAnswer == false {
		f.Print("-> ")
		choice, _ := reader.ReadString('\n')
		choice = s.Replace(choice, "\n", "", -1)
		coords := s.Split(choice, ",")

		i1, err := strconv.Atoi(coords[0])
		if err == nil { x = i1 }
		i2, err := strconv.Atoi(coords[1])
		if err == nil { y = i2 }
		if x <= 2 && x >= 0 && y <= 2 && y >= 0 {
			validAnswer = true
		} else {
			f.Println("Invalid coordinates, try again")
		}
	}
}


func game() {
	if turn % 2 == 0 {
		player = "b"
	}	else {
		player = "a"
	}
	
	choose()

	for (board[y][x] == "x" || board[y][x] == "o") { 
		f.Println("That place is already taken. Try again")
		showBoard()
		choose()
	}

	updateBoard([]int{x, y}, player)

	turn += 1
}