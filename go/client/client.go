package client

import (
	"github.com/gurrpi/daily-buyer/go/exchange"
	"github.com/gurrpi/daily-buyer/go/types"
	"github.com/gurrpi/daily-buyer/go/user"
)

type Client struct {
	User     user.User
	Exchange exchange.Exchange
}

func (c Client) Buy(order types.Order) (types.Receipt, error) {
	panic("implement me!")
}
