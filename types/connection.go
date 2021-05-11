package types

const (
	CHAT_SERVER_PORT = 3020
)

type Connection struct {
	Utente User
	IdChat int
}
