package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Compleo/compleo-server-chat/chat_connector"
	"github.com/Compleo/compleo-server-chat/types"
	"golang.org/x/net/websocket"
)

func main() {
	fmt.Println("Starting Compleo server chat")

	http.Handle("/", websocket.Handler(chat_connector.RootWSS))

	if err := http.ListenAndServe(":"+strconv.Itoa(types.CHAT_SERVER_PORT), nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
