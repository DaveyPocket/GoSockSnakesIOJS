// Game code goes here

package logic

import ("fmt"
		"os"
		"encoding/json"
		"math/rand"
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
	Body		[]points	`json:"points"`		// Slice of points
	Dir			string		`json:"direction"`	// Head direction
	//Owner		ownerT		`json:"owner:"`		// Connected user
}

type f points

type game struct {
	Snake		[]s			`json:"snakes"`		// Slice of snake
	Food		[]f			`json:"food"`		// Slice of foods
	//ClientID	int			`json:"clientId"`
//	State		stateType	`json:"state"`		// State byte
	//nowTick		byte		// Current tick byte
}

type StatePacket struct {
	game
	ClientID	int			`json:"clientId"`
}

func InitGame() (*game) {
	var g game
	g.Snake = make([]s, 0)
	g.Food = make([]f, 0)
//	g.State = empty
	return &g
}

func (g *game) AddSnake(x, y int, startSize int, startDir string) {
	b := make([]points, startSize)
	// Typecast with direction type to byte 
	// For now start direciton is always toward the right
	for i, _ := range b {
		b[i] = points{x + i, y}
	}
	nS := s{b, startDir}
	g.Snake = append(g.Snake, nS)
	fmt.Println("Added snake: ", g.Snake[len(g.Snake) - 1: len(g.Snake)])
	fmt.Println("\n\nCurrent Snakes:", g.Snake, "\n")
}

func (g *game) AddFood(x, y int) {
	nF := f{x, y}
	g.Food = append(g.Food, nF)
}

func (g *game) EatFood(tail points) {
	for i, sn := range g.Snake {
		for q, fd := range g.Food {
			if sn.Body[0] == points(fd) {
				g.Snake[i].Body = append(g.Snake[i].Body, tail)
				g.Food = append(g.Food[0:q], g.Food[q+1:]...) // Remove food
			}
		}
	}
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

func (g *game) Tick(p int64) {
	r := rand.New(rand.NewSource(p))
	var tempTail points
	for i, sn := range g.Snake {
		tempTail = g.Snake[i].Body[len(sn.Body)-1]
		for q, _ := range sn.Body[:len(sn.Body)-1] {
			g.Snake[i].Body[len(sn.Body) - q - 1] = g.Snake[i].Body[len(sn.Body) - q - 2]
		}
		switch sn.Dir {
				case "UP":
					g.Snake[i].Body[0].Y--
				case "DOWN":
					g.Snake[i].Body[0].Y++
				case "LEFT":
					g.Snake[i].Body[0].X--
				case "RIGHT":
					g.Snake[i].Body[0].X++
		}
		g.EatFood(tempTail)
		if len(g.Food) == 0 {
			g.AddFood(r.Int() % 30, r.Int() % 30)
		}
		fmt.Println(sn)
	}
}

func GetPacket(g game, clientID int) (*StatePacket) {
	sp := StatePacket{g, clientID}
	return &sp
}

func GetJSONPacket(g game, clientID int) ([]byte) {
	o, err := json.Marshal(GetPacket(g, clientID))
	if err != nil {
		panic(err)
	}
	fmt.Print("JSON packet: ")
	os.Stdout.Write(o)
	return o
}
