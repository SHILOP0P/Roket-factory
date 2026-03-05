package model

type Order struct {
	OrderUUID       string
	UserUUID        string
	PartUUIDs       []string
	TotalPrice      float64
	TransactionUUID *string
	PaymentMethod   *PaymentMethod
	Status          OrderStatus
}

type UpdateOrder struct {
	OrderUUID       string
	UserUUID        string
	PartUUIDs       *[]string
	TotalPrice      *float64
	TransactionUUID *string
	PaymentMethod   *PaymentMethod
	Status          *OrderStatus
}

type PaymentMethod int32

const (
	PaymentMethod_PAYMENT_METHOD_UNSPECIFIED    PaymentMethod = 0
	PaymentMethod_PAYMENT_METHOD_CARD           PaymentMethod = 1
	PaymentMethod_PAYMENT_METHOD_SBP            PaymentMethod = 2
	PaymentMethod_PAYMENT_METHOD_CREDIT_CARD    PaymentMethod = 3
	PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY PaymentMethod = 4
)

type OrderStatus string

const (OrderStatusPENDINGPAYMENT OrderStatus = "PENDING_PAYMENT"
	OrderStatusPAID           OrderStatus = "PAID"
	OrderStatusCANCELLED      OrderStatus = "CANCELLED"
)
	