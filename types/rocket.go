package types

import "github.com/google/uuid"

type Rocket struct {
	Id      uuid.UUID `json:"id,omitempty"`
	Name    string    `json:"name"`
	Mission string    `json:"mission"`
	Speed   int       `json:"speed"`
}
