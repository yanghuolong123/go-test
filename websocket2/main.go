package main

import (
	"./chat"
	"net/http"
)

func main() {
	server := chat.NewServer("/chat")
	go server.Listen()

	http.ListenAndServe(":8585", nil)
}
