package domain

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Hello struct {
	msg string
}

func (h *Hello) Message() string {
	return h.msg
}

func NewMessageForUser(user string) (*Hello, error) {
	msgs := [3]string{"Hello %s world", "Aloha %", "Glad to see you %s"}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(msgs))))
	if err != nil {
		return nil, err
	}

	return &Hello{msg: fmt.Sprintf(msgs[n.Int64()], user)}, nil
}
