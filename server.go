package main

import (
    "log"
    "net/http"
	"fmt"
    "github.com/googollee/go-socket.io"
	"time"
	"./game"
)

func main() {
	g := logic.InitGame()
	g.AddSnake(1, 1, 5, "LEFT")
	g.AddSnake(2, 3, 8, "RIGHT")
    server, err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }
    server.On("connection", func(so socketio.Socket) {
	
        log.Println("on connection")
        so.Join("main")
        so.On("join game", func() {
            log.Println("Received join game")
            		
        })
		log.Println("After join game...")
		so.BroadcastTo("main", "init setup", g)

		so.On("ready", func() {
			log.Println("Received ready")
			go func() {
			for {
				log.Println("tick")
				so.BroadcastTo("main", "tick")
				time.Sleep(500 * time.Millisecond)
			}}()
		})
        so.On("disconnection", func() {
            log.Println("on disconnect")
        })
    })
    server.On("error", func(so socketio.Socket, err error) {
        log.Println("error:", err)
    })

    http.Handle("/socket.io/", server)
    http.Handle("/", http.FileServer(http.Dir("./asset")))
	fmt.Println(http.Dir("./asset"))
    log.Println("Serving at localhost:5000...")
    log.Fatal(http.ListenAndServe(":5000", nil))
}

func servTick(so socketio.Socket) {
	for {
		so.BroadcastTo("main", "tick")
		time.Sleep(500 * time.Millisecond)
	}

}
