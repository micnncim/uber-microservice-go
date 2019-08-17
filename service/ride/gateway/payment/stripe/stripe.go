package stripe

import (
	"github.com/micnncim/uber-microservice-go/service/ride/gateway/payment"
)

type Stripe struct {
}

func NewStripe() payment.Payment {
	return &Stripe{}
}
