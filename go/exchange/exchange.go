package exchange

import (
	"github.com/gurrpi/daily-buyer/go/types"
	"github.com/gurrpi/daily-buyer/go/user"
)

type Exchange interface {
	Account(user user.User) (types.ResponseAccount, error)
	Order(user user.User, order types.Order) (types.ResponseOrder, error)
}
