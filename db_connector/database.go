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
