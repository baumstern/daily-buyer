package user

import (
	"github.com/gurrpi/daily-buyer/go/types"
)

type Buyer interface {
	Buy(order types.Order) (types.Receipt, error)
}

type User struct {
	Credential types.Credential
}
