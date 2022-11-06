package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		fmt.Println("connected:", so.ID())
		return nil
	})

	server.OnEvent("/halo", "notice", func(so socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		so.Emit("reply", "have "+msg)
	})

	server.OnDisconnect("/", func(c socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket/io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("serving localhost on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
