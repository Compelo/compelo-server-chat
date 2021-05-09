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
			//STRUTTURA: [idChat] [idMittente] [idDestinatario] [tipo] [contenuto]
			_sessione, e := sessione.FindConnection(addr)
			if e == "Non trovato nulla" {
				fmt.Println("Errore -> Non trovato la sessione per ", addr)
				return
			}

			//Prendo i dati
			idChat, err := strconv.ParseInt(arr[1], 10, 64)
			if err != nil {
				fmt.Println("Errore: ", err)
			}

			idMittente, err := strconv.ParseInt(arr[2], 10, 64)
			if err != nil {
				fmt.Println("Errore: ", err)
			}

			idDestinatario, err := strconv.ParseInt(arr[3], 10, 64)
			if err != nil {
				fmt.Println("Errore: ", err)
			}

			tipo, err := strconv.ParseInt(arr[4], 10, 64)
			if err != nil {
				fmt.Println("Errore: ", err)
			}

			contenuto := arr[5]
			if err != nil {
				fmt.Println("Errore: ", err)
			}

			db_connector.NewMessage(sessione.Connessioni[_sessione], int(idMittente), int(idDestinatario), int(tipo), contenuto, int(idChat))
		case "REMMESSAGE":
			//Elimina il messaggio
			//STRUTTURA: [idMessaggio]
			idMessaggio, err := strconv.ParseInt(arr[1], 10, 64)
			if err != nil {
				fmt.Println("Errore: ", err)
			}

			db_connector.EliminaMessaggio(int(idMessaggio))
		case "GETMESSAGES":
			//Prendi tutti i messaggi di una chat
			//STRUTTURA: [idChat]
		default:
			fmt.Println(ws.Request().RemoteAddr, " -> NOT PERMITTED")
		}
	}

	db_connector.Disconnetti()
}
