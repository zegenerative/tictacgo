package main

import "fmt"
 
var ends = [2][3]string{
	{"o","o","o"},
	{"x","x","x"},
}

func checkRows(board [3][3]string) {
	fmt.Println("check:", board[0])
	// horizontal1 := [1][3]string{
	// 	{board[0][0]},
	// 	{board[0][1]},
	// 	{board[0][2]},
	// }
	// horizontal2 := board[1][0] + board[1][1] + board[1][2]
	// horizontal3 := board[2][0] + board[2][1] + board[2][2]

	// vertical1 := board[0][0] + board[1][0] + board[2][0]
	// vertical2 := board[0][1] + board[1][1] + board[2][1]
	// vertical3 := board[0][2] + board[1][2] + board[2][2]

	// diagonal1 := board[0][0] + board[1][1] + board[2][2]
	// diagonal2 := board[2][0] + board[1][1] + board[0][2]
	// fmt.Println(horizontal1) 
		// horizontal2, horizontal3, vertical1,vertical2, vertical3, diagonal1, diagonal2)
}



