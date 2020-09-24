package main

import (
	f "fmt"
)

var board [][]string

func updateBoard(choice [][]int) {
	f.Println(choice)
}

func showBoard() {
	updateBoard([][]int{
		{1, 0, 0},{},{},
	})
}