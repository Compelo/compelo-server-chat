package db_connector

import (
	"database/sql"

	"github.com/Compleo/compleo-server-chat/types"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func Connetti() {
	//Connessione al database:
	var sqlError error

	Database, sqlError = sql.Open("sqlite3", "./data.db")
	if sqlError != nil {
		panic(sqlError)
	}
}

func Disconnetti() {
	Database.Close()
}

func NewMessage(sessione types.Connection, mittente int, destinatario int, tipo int, contenuto string, chat int) {
	//TODO: IMPLEMENTA
}

func EliminaMessaggio(id int) {
	//TODO: IMPLEMENTA
}

func GetMessagesFromIDChat(id int) string {
	//TODO: IMPLEMENTA
	return ""
}
