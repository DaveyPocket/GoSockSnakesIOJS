// Game code goes here

package logic

import ("fmt"
		"os"
		"encoding/json"
		)

const (
		sizeX = 60
		sizeY = 60
		)

type stateType byte

const (
		empty stateType = iota
		running
		pause
	)

type points struct {
	X	int			`json:"x"`
	Y	int			`json:"y"`					// Coordinates
}

type ownerT	byte

type s struct {
	Body		[]points	`json:"body"`		// Slice of points
	Dir			byte		`json:"dir"`		// Head direction
	Owner		ownerT		`json:"owner:"`		// Connected user
}

type game struct {
	Snake		[]s			`json:"snakes"`		// Slice of snakes
//	food		[]f			`json:"food"`		// Slice of foods
	State		stateType	`json:"state"`		// State byte
	//nowTick		byte		// Current tick byte
}

func InitGame() (*game) {
	var g game
	g.Snake = make([]s, 0)
//	g.food = make([]food, 0)
	g.State = empty
	return &g
}

func (g *game) AddSnake(x, y int, startSize int, startDir byte, owner ownerT) {
	b := make([]points, startSize)
	// Typecast with direction type to byte 
	// For now start direciton is always toward the right
	for i, _ := range b {
		b[i] = points{x + i, y}
	}
	nS := s{b, startDir, owner}
	g.Snake = append(g.Snake, nS)
	fmt.Println("Added snake: ", g.Snake[len(g.Snake) - 1: len(g.Snake)])
}

func (g *game) GetJSON() ([]byte) {
	o, err := json.Marshal(*g)
	if err != nil {
		panic(err)
	}
	fmt.Print("JSON: ")
	os.Stdout.Write(o)
	return o
}
