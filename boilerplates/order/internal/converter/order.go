package converter

import (
	"fmt"
	"order/internal/model"

	"github.com/google/uuid"

	orderv1 "shared/pkg/openapi/order/v1"
)

func CreateOrderRequestToModel(req *orderv1.CreateOrderRequest) model.Order {
	if req == nil {
		return model.Order{}
	}

	return model.Order{
		UserUUID:  req.GetUserUUID().String(),
		PartUUIDs: UuidsToStrings(req.GetPartUuids()),
	}
}

func GetOrderByUUIDParamsToString(params orderv1.GetOrderByUUIDParams) string {
	return params.OrderUUID.String()
}

func PayOrderRequestToModel(params orderv1.PayOrderParams, req *orderv1.PayOrderRequest) model.UpdateOrder {
	if req == nil {
		return model.UpdateOrder{OrderUUID: params.OrderUUID.String()}
	}

	paymentMethod := PaymentMethodFromAPI(req.GetPaymentMethod())

	return model.UpdateOrder{
		OrderUUID:     params.OrderUUID.String(),
		PaymentMethod: &paymentMethod,
	}
}

func CancelOrderParamsToModel(params orderv1.CancelOrderParams) model.UpdateOrder {
	status := model.OrderStatusCANCELLED
	return model.UpdateOrder{
		OrderUUID: params.OrderUUID.String(),
		Status:    &status,
	}
}

func OrderToGetOrderByUUIDRes(in model.Order) (*orderv1.OrderDto, error) {
	orderUUID, err := uuid.Parse(in.OrderUUID)
	if err != nil {
		return nil, fmt.Errorf("parse order uuid: %w", err)
	}

	userUUID, err := uuid.Parse(in.UserUUID)
	if err != nil {
		return nil, fmt.Errorf("parse user uuid: %w", err)
	}

	partUUIDs, err := StringsToUUIDs(in.PartUUIDs)
	if err != nil {
		return nil, fmt.Errorf("parse part uuids: %w", err)
	}

	dto := &orderv1.OrderDto{
		OrderUUID:  orderUUID,
		UserUUID:   userUUID,
		PartUuids:  partUUIDs,
		TotalPrice: in.TotalPrice,
		Status:     OrderStatusToAPI(in.Status),
	}

	if in.TransactionUUID != nil {
		txUUID, parseErr := uuid.Parse(*in.TransactionUUID)
		if parseErr != nil {
			return nil, fmt.Errorf("parse transaction uuid: %w", parseErr)
		}
		dto.TransactionUUID = orderv1.NewOptNilUUID(txUUID)
	}

	if in.PaymentMethod != nil {
		dto.PaymentMethod = orderv1.NewOptPaymentMethod(PaymentMethodToAPI(*in.PaymentMethod))
	}

	return dto, nil
}

func OrderToCreateOrderResponse(in model.Order) (*orderv1.CreateOrderResponse, error) {
	orderUUID, err := uuid.Parse(in.OrderUUID)
	if err != nil {
		return nil, fmt.Errorf("parse order uuid: %w", err)
	}

	return &orderv1.CreateOrderResponse{
		OrderUUID:  orderUUID,
		TotalPrice: in.TotalPrice,
	}, nil
}

func OrderToPayOrderResponse(in model.Order) (*orderv1.PayOrderResponse, error) {
	if in.TransactionUUID == nil {
		return nil, fmt.Errorf("transaction uuid is empty")
	}

	txUUID, err := uuid.Parse(*in.TransactionUUID)
	if err != nil {
		return nil, fmt.Errorf("parse transaction uuid: %w", err)
	}

	return &orderv1.PayOrderResponse{TransactionUUID: txUUID}, nil
}

func PaymentMethodFromAPI(in orderv1.PaymentMethod) model.PaymentMethod {
	switch in {
	case orderv1.PaymentMethodPAYMENTMETHODCARD:
		return model.PaymentMethod_PAYMENT_METHOD_CARD
	case orderv1.PaymentMethodPAYMENTMETHODSBP:
		return model.PaymentMethod_PAYMENT_METHOD_SBP
	case orderv1.PaymentMethodPAYMENTMETHODCREDITCARD:
		return model.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD
	case orderv1.PaymentMethodPAYMENTMETHODINVESTORMONEY:
		return model.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY
	default:
		return model.PaymentMethod_PAYMENT_METHOD_UNSPECIFIED
	}
}

func PaymentMethodToAPI(in model.PaymentMethod) orderv1.PaymentMethod {
	switch in {
	case model.PaymentMethod_PAYMENT_METHOD_CARD:
		return orderv1.PaymentMethodPAYMENTMETHODCARD
	case model.PaymentMethod_PAYMENT_METHOD_SBP:
		return orderv1.PaymentMethodPAYMENTMETHODSBP
	case model.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD:
		return orderv1.PaymentMethodPAYMENTMETHODCREDITCARD
	case model.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY:
		return orderv1.PaymentMethodPAYMENTMETHODINVESTORMONEY
	default:
		return orderv1.PaymentMethodPAYMENTMETHODUNSPECIFIED
	}
}

func OrderStatusFromAPI(in orderv1.OrderStatus) model.OrderStatus {
	switch in {
	case orderv1.OrderStatusPAID:
		return model.OrderStatusPAID
	case orderv1.OrderStatusCANCELLED:
		return model.OrderStatusCANCELLED
	default:
		return model.OrderStatusPENDINGPAYMENT
	}
}

func OrderStatusToAPI(in model.OrderStatus) orderv1.OrderStatus {
	switch in {
	case model.OrderStatusPAID:
		return orderv1.OrderStatusPAID
	case model.OrderStatusCANCELLED:
		return orderv1.OrderStatusCANCELLED
	default:
		return orderv1.OrderStatusPENDINGPAYMENT
	}
}

func UuidsToStrings(in []uuid.UUID) []string {
	if len(in) == 0 {
		return nil
	}

	out := make([]string, 0, len(in))
	for _, id := range in {
		out = append(out, id.String())
	}

	return out
}

func StringsToUUIDs(in []string) ([]uuid.UUID, error) {
	if len(in) == 0 {
		return nil, nil
	}

	out := make([]uuid.UUID, 0, len(in))
	for _, s := range in {
		id, err := uuid.Parse(s)
		if err != nil {
			return nil, err
		}
		out = append(out, id)
	}

	return out, nil
}

