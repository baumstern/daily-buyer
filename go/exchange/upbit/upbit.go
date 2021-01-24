package upbit

import (
	"github.com/gurrpi/daily-buyer/go/types"
	"github.com/gurrpi/daily-buyer/go/user"
)

type Upbit struct {
}

func (u Upbit) Account(user user.User) (types.ResponseAccount, error) {
	panic("implement me")
}

func (u Upbit) Order(user user.User, order types.Order) (types.ResponseOrder, error) {
	panic("implement me")
}
