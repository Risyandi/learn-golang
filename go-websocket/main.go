package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)

	// ** not broadcast to all connected
	// server.OnConnect("/", func(so socketio.Conn) error {
	// 	so.SetContext("")
	// 	fmt.Println("connected:", so.ID())
	// 	return nil
	// })

	// ** broadcast to all connected
	server.OnConnect("/chat", func(so socketio.Conn) error {
		so.SetContext("")
		fmt.Println("chat user connected: ", so.ID())
		so.Join("room1")
		return nil
	})

	// server event for socket.io
	server.OnEvent("/chat", "reply", func(so socketio.Conn, messageText string) {
		so.Emit("reply", messageText)
	})

	server.OnEvent("/chat", "messageText", func(so socketio.Conn, messageText string) string {
		nameSpace := so.Namespace()
		server.BroadcastToRoom(nameSpace, "room1", "reply", messageText)
		return messageText
	})

	server.OnEvent("/", "bye", func(so socketio.Conn) string {
		last := so.Context().(string)
		so.Emit("bye", last)
		so.Close()
		return last
	})

	server.OnError("/chat", func(so socketio.Conn, message error) {
		fmt.Println("error: ", message)
	})

	server.OnDisconnect("/chat", func(so socketio.Conn, reason string) {
		fmt.Println("closed: ", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("serving localhost on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
