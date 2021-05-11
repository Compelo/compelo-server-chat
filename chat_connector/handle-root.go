package chat_connector

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Compleo/compleo-server-chat/db_connector"
	"github.com/Compleo/compleo-server-chat/types"
	"golang.org/x/net/websocket"
)

var sessione types.Session

func RootWSS(ws *websocket.Conn) {
	var err error

	db_connector.Connetti()

	for {
		var recived string

		if err = websocket.Message.Receive(ws, &recived); err != nil {
			fmt.Println("Errore: ", err)
			break
		}

		arr := strings.Split(recived, " ")
		addr := ws.Request().RemoteAddr

		switch arr[0] {
		case "REGME":
			//Registro una nuova connessione
			id, err := strconv.ParseInt(arr[1], 10, 64)
			if err != nil {
				fmt.Println("Errore: ", err)
			}

			sessione.NewConenction(addr, id)
			fmt.Println(addr, " -> Aggiunto nella sessione")
		case "REME":
			//Elimina la sessione
			sessione.RemoveConnection(addr)
			fmt.Println(addr, " -> Rimosso dalla sessione")
		case "ADDMESSAGE":
			//Nuovo messaggio
			//STRUTTURA: [contenuto]
			_sessione, e := sessione.FindConnection(addr)
			if e == "Non trovato nulla" {
				fmt.Println("Errore -> Non trovato la sessione per ", addr)
				return
			}

			contenuto := arr[1]
			if err != nil {
				fmt.Println("Errore: ", err)
			}

			db_connector.NewMessage(sessione.Connessioni[_sessione], contenuto, sessione.Connessioni[_sessione].IdChat)

			fmt.Println(addr, " -> Aggiunto il messaggio")
		case "GETMESSAGES":
			//Prendi tutti i messaggi di una chat
			//STRUTTURA: [idChat]
			idChat, err := strconv.Atoi(arr[1])
			if err != nil {
				fmt.Println("Errore: ", err)
			}

			msg := db_connector.GetMessagesFromIDChat(idChat)

			ws.Write([]byte(msg))

			fmt.Println(addr, " -> Ha richiesto i messaggi")
		case "CHANGECHAT":
			//Prendi tutti i messaggi di una chat
			//STRUTTURA: [idChat]
			i, err := strconv.ParseInt(arr[1], 10, 64)
			if err != nil {
				fmt.Println("Errore: ", err)
			}

			pos, _ := sessione.FindConnection(addr)
			sessione.Connessioni[pos].IdChat = int(i)

			fmt.Println(addr, " -> Ha cambiato la chat corrente")
		default:
			fmt.Println(ws.Request().RemoteAddr, " -> NOT PERMITTED")
		}
	}

	db_connector.Disconnetti()
}
