package types

import "time"

const (
	TYPE_TEXT  = 0
	TYPE_IMAGE = 1
)

const (
	FROM_WHO_A = 0
	FROM_WHO_B = 1
)

type Message struct {
	Type        int
	CharStream  string
	FromWho		int
	CreatedTime time.Time
}
