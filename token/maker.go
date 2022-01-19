package token

import (
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}
