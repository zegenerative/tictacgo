package main

import "fmt"
 
var ends = [2][3]string{
	{"o","o","o"},
	{"x","x","x"},
}

func checkRows(board [3][3]string) {
	horizontal1 := [3]string{
		board[2][0], board[2][1], board[2][2],
	}
	horizontal2 := [3]string{
		board[1][0], board[1][1], board[1][2],
	}
	horizontal3 := [3]string{
		board[0][0], board[0][1], board[0][2],
	}

	vertical1 := [3]string{
		board[2][0], board[1][0], board[0][0],
	}
	vertical2 := [3]string{
		board[2][1], board[1][1], board[0][1],
	}
	vertical3 := [3]string{
		board[2][2], board[1][2], board[0][2],
	}

	diagonal1 := [3]string{
		board[2][0], board[1][1], board[0][2],
	}
	diagonal2 := [3]string{
		board[0][0], board[1][1], board[2][2],
	}

	if (
		(horizontal1 == ends[0] || horizontal1 == ends[1]) ||
		(horizontal2 == ends[0] || horizontal2 == ends[1]) ||
		(horizontal3 == ends[0] || horizontal1 == ends[1]) ||
		(vertical1 == ends[0] || vertical1 == ends[1]) ||
		(vertical2 == ends[0] || vertical2 == ends[1]) ||
		(vertical3 == ends[0] || vertical3 == ends[1]) ||
		(diagonal1 == ends[0] || diagonal1 == ends[1]) ||
		(diagonal2 == ends[0] || diagonal2 == ends[1])) {
		gameOver = true
	}
	if gameOver == true {
		fmt.Println("Game Over. Player " + player + " wins")
	}
}



