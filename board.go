package main

import (
	f "fmt"
)

var board [3][3]string

func updateBoard(choice []int, player string) {
	var x = choice[0]
	var y = choice[1]
	if player == "a" {
		board[y][x] = "x"
	} else {
		board[y][x] = "o"
	}
}

func showBoard() {
	f.Println(board[2])
	f.Println(board[1])
	f.Println(board[0])
}