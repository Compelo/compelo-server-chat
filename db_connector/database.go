package db_connector

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/Compleo/compleo-server-chat/types"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var DatabaseSQLIte *sql.DB
var DatabaseMySQL *sql.DB

func Connetti() {
	//Connessione al database:
	var sqlError error

	DatabaseSQLIte, sqlError = sql.Open("sqlite3", "./data.db")
	if sqlError != nil {
		panic(sqlError)
	}

	DatabaseMySQL, sqlError = sql.Open("mysql", sqlVal)
	if sqlError != nil {
		panic(sqlError.Error())
	}
}

func Disconnetti() {
	DatabaseSQLIte.Close()
	DatabaseMySQL.Close()
}

func NewMessage(sessione types.Connection, contenuto string, chat int) {
	//TODO: IMPLEMENTA
	//idChat (v)
	//tipo -> 0
	//content = contenuto
	//idUtenteMittente = sessione.Utente.DatabaseID
	//idUtenteDestinatario = Eseguo una query per trovare la chat nel database principale
	//						 Dai risultati prendo i due ip degli utenti e li confronto con quello che ho già

	var idUtenteDestinatario int

	queyRes, queyErr := DatabaseMySQL.Query("SELECT * FROM `chat` WHERE `ID`='" + fmt.Sprint(chat) + "'")
	if queyErr != nil {
		fmt.Println(queyErr)
		return
	}

	var toReturn types.Chat
	for queyRes.Next() {
		scanErr := queyRes.Scan(&toReturn.ID, &toReturn.IDUtenteRichiedente, &toReturn.IDUtenteDestinatario)
		if scanErr != nil {
			fmt.Println(scanErr)
			return
		}
	}

	fmt.Println(toReturn)

	if toReturn.IDUtenteRichiedente == sessione.Utente.DatabaseID {
		idUtenteDestinatario = int(toReturn.IDUtenteDestinatario)
	} else if toReturn.IDUtenteDestinatario == sessione.Utente.DatabaseID {
		idUtenteDestinatario = int(toReturn.IDUtenteRichiedente)
	} else {
		fmt.Println("Errore")
	}

	//Popolo l'oggetto da inserire
	var messaggio types.Message
	messaggio.Content = contenuto
	messaggio.Tipo = 0
	messaggio.IDChat = sessione.IdChat
	messaggio.IDUtenteDestinatario = idUtenteDestinatario
	messaggio.IDUtenteMittente = int(sessione.Utente.DatabaseID)

	query := "INSERT INTO `messaggi`(`IDChat`, `Tipo`, `Content`, `IDUtenteMittente`, `IDUtenteDestinatario`) VALUES ('" + fmt.Sprint(messaggio.IDChat) + "', '" + fmt.Sprint(messaggio.Tipo) + "', '" + messaggio.Content + "', '" + fmt.Sprint(messaggio.IDUtenteMittente) + "', '" + fmt.Sprint(messaggio.IDUtenteDestinatario) + "')"
	DatabaseSQLIte.Prepare(query)
	_, err := DatabaseSQLIte.Exec(query)
	if err != nil {
		fmt.Println(queyErr)
		return
	}

}

func GetMessagesFromIDChat(id int) string {
	queyRes, queyErr := DatabaseSQLIte.Query("SELECT * FROM `messaggi` WHERE `IDChat`=" + fmt.Sprint(id) + "")
	if queyErr != nil {
		fmt.Println(queyErr)
		return ""
	}

	var toReturn []types.Message

	for queyRes.Next() {
		var g types.Message
		scanErr := queyRes.Scan(&g.ID, &g.IDChat, &g.Tipo, &g.Content, &g.IDUtenteMittente, &g.IDUtenteDestinatario)
		if scanErr != nil {
			fmt.Println(scanErr)
			return ""
		}

		toReturn = append(toReturn, g)
	}

	//Create JSON
	j, _ := json.Marshal(toReturn)

	return string(j)
}
