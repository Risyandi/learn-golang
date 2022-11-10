package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	// initialization the socket server
	server := socketio.NewServer(nil)

	// handling socket event connected
	// and set default not join in others room, when connected and send message not broadcast to all connected
	// server.OnConnect("/", func(so socketio.Conn) error {
	// 	so.SetContext("")
	// 	fmt.Println("connected:", so.ID())
	// 	return nil
	// })

	// handling socket event connected
	// and default join in room with name "room" for broadcast to all connected with name room same
	server.OnConnect("/chat", func(so socketio.Conn) error {
		// set default context is empty and then join to room with name "room1"
		so.SetContext("")

		// print on console log server
		fmt.Println("id users connected: ", so.ID())

		so.Join("room1")
		return nil
	})

	// handling socket event call nameEvent "reply"
	server.OnEvent("/chat", "reply", func(so socketio.Conn, messageText string) {
		// test case
		// checkURL := so.URL()
		// fmt.Println("url: ", checkURL)

		// checkLocalAddr := so.LocalAddr()
		// fmt.Println("LocalAddr: ", checkLocalAddr)

		// not test case
		fmt.Println("messageText: ", messageText)
		so.Emit("reply", messageText)
		// return messageText
	})

	// handling socket event call nameEvent "typing"
	server.OnEvent("/chat", "typing", func(so socketio.Conn, messageText string) string {
		nameSpace := so.Namespace()
		server.BroadcastToRoom(nameSpace, "room1", "typing", messageText)
		return messageText
	})

	// handling socket event error
	server.OnError("/chat", func(so socketio.Conn, message error) {
		// print on console log server
		// messageError := message
		// fmt.Println("error: ", messageError)
		fmt.Println("error: ", message)
	})

	// handling socket event disconnected
	server.OnDisconnect("/chat", func(so socketio.Conn, reason string) {
		// print on console log server
		// messageReason := reason
		// fmt.Println("closed: ", messageReason)
		fmt.Println("closed: ", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("serving localhost on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
