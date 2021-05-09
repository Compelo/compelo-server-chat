package chat_connector

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Compleo/compleo-server-chat/types"
	"golang.org/x/net/websocket"
)

var sessione types.Session

func RootWSS(ws *websocket.Conn) {
	var err error

	for {
		var recived string

		if err = websocket.Message.Receive(ws, &recived); err != nil {
			fmt.Println("Can't receive")
			fmt.Println(err)
			break
		}

		arr := strings.Split(recived, " ")

		switch arr[0] {
		case "REGME":
			//Registro una nuova connessione
			addr := ws.Request().RemoteAddr
			id, err := strconv.ParseInt(arr[1], 10, 64)
			if err != nil {
				fmt.Println("Errore: ", err)
			}

			sessione.NewConenction(addr, id)
			fmt.Println(addr, " -> Aggiunto come sessione")
		case "REME":
			addr := ws.Request().RemoteAddr

			sessione.RemoveConnection(addr)
			fmt.Println(addr, " -> Rimosso")
		}
	}
}
