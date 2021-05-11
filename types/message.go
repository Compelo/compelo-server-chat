package types

const (
	TYPE_TEXT  = 0
	TYPE_IMAGE = 1
)

const (
	FROM_WHO_A = 0
	FROM_WHO_B = 1
)

type Message struct {
	ID                   int    `json:"id"`
	IDChat               int    `json:"idChat"`
	Tipo                 int    `json:"tipo"`
	Content              string `json:"contenuto"`
	IDUtenteMittente     int    `json:"idUtenteMittente"`
	IDUtenteDestinatario int    `json:"idUtenteDestinatario"`
}
