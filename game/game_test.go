package logic

import ("testing"
		"fmt"
		"encoding/json"
		)

func TestInitGame(t *testing.T) {
	g := InitGame()
	fmt.Println(*g)
}

// Remaining functions should not run if TestInitGame fails
func TestAddSnake(t *testing.T) {
	g := InitGame()
	g.AddSnake(1, 2, 5, "RIGHT")
	fmt.Println(g.Snake[0].Body)
}

func TestGetJSON(t *testing.T) {
	g := InitGame()
	g.AddSnake(1, 2, 5, "RIGHT")
	g.AddSnake(2, 3, 1, "RIGHT")
	q := g.GetJSON()
	ung := InitGame()
	err := json.Unmarshal(q, &ung)
	if err != nil {
		panic(err)
	}
	fmt.Println("\n", *ung)
}


