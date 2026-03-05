package converter

import (
	"order/internal/model"
	repoModel "order/internal/repository/model"
)

func OrderToRepoModel(in model.Order) repoModel.Order {
	return repoModel.Order{
		OrderUUID:       in.OrderUUID,
		UserUUID:        in.UserUUID,
		PartUUIDs:       append([]string(nil), in.PartUUIDs...),
		TotalPrice:      in.TotalPrice,
		TransactionUUID: cloneStringPtr(in.TransactionUUID),
		PaymentMethod:   paymentMethodToRepo(in.PaymentMethod),
		Status:          repoModel.OrderStatus(in.Status),
	}
}

func OrderToModel(in repoModel.Order) model.Order {
	return model.Order{
		OrderUUID:       in.OrderUUID,
		UserUUID:        in.UserUUID,
		PartUUIDs:       append([]string(nil), in.PartUUIDs...),
		TotalPrice:      in.TotalPrice,
		TransactionUUID: cloneStringPtr(in.TransactionUUID),
		PaymentMethod:   paymentMethodToModel(in.PaymentMethod),
		Status:          model.OrderStatus(in.Status),
	}
}

func OrderUpdateToRepoModel(in model.UpdateOrder) repoModel.UpdateOrder {
	return repoModel.UpdateOrder{
		OrderUUID:       in.OrderUUID,
		UserUUID:        in.UserUUID,
		PartUUIDs:       cloneStringSlicePtr(in.PartUUIDs),
		TotalPrice:      cloneFloat64Ptr(in.TotalPrice),
		TransactionUUID: cloneStringPtr(in.TransactionUUID),
		PaymentMethod:   paymentMethodToRepo(in.PaymentMethod),
		Status:          orderStatusToRepo(in.Status),
	}
}

func OrderUpdateToModel(in repoModel.UpdateOrder) model.UpdateOrder {
	return model.UpdateOrder{
		OrderUUID:       in.OrderUUID,
		UserUUID:        in.UserUUID,
		PartUUIDs:       cloneStringSlicePtr(in.PartUUIDs),
		TotalPrice:      cloneFloat64Ptr(in.TotalPrice),
		TransactionUUID: cloneStringPtr(in.TransactionUUID),
		PaymentMethod:   paymentMethodToModel(in.PaymentMethod),
		Status:          orderStatusToModel(in.Status),
	}
}

func paymentMethodToRepo(in *model.PaymentMethod) *repoModel.PaymentMethod {
	if in == nil {
		return nil
	}
	v := repoModel.PaymentMethod(*in)
	return &v
}

func paymentMethodToModel(in *repoModel.PaymentMethod) *model.PaymentMethod {
	if in == nil {
		return nil
	}
	v := model.PaymentMethod(*in)
	return &v
}

func orderStatusToRepo(in *model.OrderStatus) *repoModel.OrderStatus {
	if in == nil {
		return nil
	}
	v := repoModel.OrderStatus(*in)
	return &v
}

func orderStatusToModel(in *repoModel.OrderStatus) *model.OrderStatus {
	if in == nil {
		return nil
	}
	v := model.OrderStatus(*in)
	return &v
}

func cloneStringSlicePtr(in *[]string) *[]string {
	if in == nil {
		return nil
	}
	out := append([]string(nil), (*in)...)
	return &out
}

func cloneFloat64Ptr(in *float64) *float64 {
	if in == nil {
		return nil
	}
	v := *in
	return &v
}

func cloneStringPtr(in *string) *string {
	if in == nil {
		return nil
	}
	v := *in
	return &v
}
