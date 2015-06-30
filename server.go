package main

import (
    "log"
    "net/http"
	"fmt"
    "github.com/googollee/go-socket.io"
	"time"
)

func main() {
    server, err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }
    server.On("connection", func(so socketio.Socket) {
		//go servTick(so)
        log.Println("on connection")
        so.Join("main")
        so.On("join game", func(msg string) {
            //log.Println("emit:", so.Emit("chat message", msg))
            so.BroadcastTo("main", "init setup", "test")
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
