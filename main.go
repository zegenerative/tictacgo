package main

var gameOver = false

func main() {
	for !gameOver {
		gameStatus()
		showBoard()
		game()
	}
}