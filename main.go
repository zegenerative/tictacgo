package main

var gameOver = false
var player string

func main() {
	for !gameOver {
		showBoard()
		game()
		gameStatus()
	}
	showBoard()
}