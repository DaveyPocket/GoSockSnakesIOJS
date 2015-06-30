package logic

import ("testing"
		"fmt"
		)

func TestInitGame(t *testing.T) {
	g := InitGame()
	fmt.Println(*g)
}

// Remaining functions should not run if TestInitGame fails
func TestAddSnake(t *testing.T) {
	g := InitGame()
	g.AddSnake(1, 2, 5, 0, ownerT(0))
	fmt.Println(g.Snake[0].Body)
}

func TestGetJSON(t *testing.T) {
	g := InitGame()
	g.AddSnake(1, 2, 5, 0, ownerT(0))
	g.AddSnake(2, 3, 1, 0, ownerT(0))
	g.GetJSON()
}

