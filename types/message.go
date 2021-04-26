package types

import "time"

const (
	TYPE_TEXT  = 0
	TYPE_IMAGE = 1
)

type Message struct {
	Type        int
	CharStream  string
	CreatedTime time.Time
}
