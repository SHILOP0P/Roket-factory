package model

import (
	"errors"
)

var ErrOrderNotFound = errors.New("Order not found")
var ErrFailCreated = errors.New("Order not created")
var ErrFailPayed = errors.New("Order not payed")
var ErrFailCancel = errors.New("Order not canceled")