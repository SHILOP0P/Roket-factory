package converter

import (
	model "payment/internal/model"
	paymentV1 "shared/pkg/proto/payment/v1"
)

func PaymentMethodProtoToModel(in paymentV1.PaymentMethod) model.PaymentMethod {
	switch in {
	case paymentV1.PaymentMethod_PAYMENT_METHOD_CARD:
		return model.PaymentMethod_PAYMENT_METHOD_CARD
	case paymentV1.PaymentMethod_PAYMENT_METHOD_SBP:
		return model.PaymentMethod_PAYMENT_METHOD_SBP
	case paymentV1.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD:
		return model.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD
	case paymentV1.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY:
		return model.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY
	default:
		return model.PaymentMethod_PAYMENT_METHOD_UNSPECIFIED
	}
}
