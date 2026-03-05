package model

import (
	"errors"
)

var ErrOrderNotFound = errors.New("Order not found")
var ErrFailCreated = errors.New("Order not created")