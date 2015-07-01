package main

import (
    "log"
    "net/http"
	"fmt"
    "github.com/googollee/go-socket.io"
	"time"
	"./game"
)
type mserve struct {
	list	[]users
}
type users struct {
	sock			socketio.Socket
	clientID		int
}

func main() {
	var pause bool
	clients := 0
	m := new(mserve)
	m.list = make([]users, 10)
	g := logic.InitGame()
    server, err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }
	
	// Make functions (Could this be done differently???)
	readyFunc := func (msg string) {
		log.Println("Received ready", msg)
		pause = false
/*		go func() {
			for {
				log.Println("tick")
				so.Emit("tick")
				time.Sleep(500 * time.Millisecond)
			}
		}(v)*/
	}


	connFunc := func(so socketio.Socket) {	
        log.Println("Client connected")
		pause = true
        so.Join("main")
		m.list[clients] = users{so, clients}
		clients++
		log.Println(&m.list)
        so.On("join game", func(msg string) {
			g.AddSnake(20, 20+clients, 5, "LEFT")
            log.Println("Received join game", msg)
			sp := logic.GetPacket(*g, clients - 1)
			so.Emit("init setup", sp)
			so.BroadcastTo("main", "init setup", sp)
			if clients == 1 {
				log.Println("Start ticker")
				go func() {
					for {
						if pause == false {
							log.Println("tick")
							so.Emit("tick")
							so.BroadcastTo("main", "tick")
							g.Tick()
						}
						time.Sleep(200 * time.Millisecond)
					}
				}()
			}
        })
		so.On("ready", readyFunc)
		so.On("change", func(cmd string){
			sp := logic.GetPacket(*g, clients - 1)
			g.Snake[m.getId(so)].Dir = cmd
			so.Emit("init setup", sp)
			so.BroadcastTo("main", "init setup", sp)

		})
		so.On("disconnection", disConnFunc)
	}

	server.On("connection", connFunc)
	server.On("error", func(so socketio.Socket, err error) {
        log.Println("error:", err)
    })

    http.Handle("/socket.io/", server)
    http.Handle("/", http.FileServer(http.Dir("./asset")))
	fmt.Println(http.Dir("./asset"))
    log.Println("Serving at localhost:5000...")
    log.Fatal(http.ListenAndServe(":5000", nil))
}

func disConnFunc(so socketio.Socket) {
	log.Println("on disconnect", so)
}

func (m* mserve) getId(sock socketio.Socket) (int) {
	for _, u := range m.list {
		fmt.Println("\n", sock, "\n", u.sock)
		if u.sock == sock {
			fmt.Println("true")
			return u.clientID
		}
	} 
	fmt.Println("False")
	return 0
}
