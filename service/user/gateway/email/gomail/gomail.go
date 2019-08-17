package gomail

import (
	"github.com/micnncim/uber-microservice-go/service/user/gateway/email"
)

type Gomail struct {
}

func NewGomail() email.Email {
	return &Gomail{}
}
