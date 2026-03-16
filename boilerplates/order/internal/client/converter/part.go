package converter

import (
	"time"

	"order/internal/model"
	inventoryV1 "shared/pkg/proto/inventory/v1"
	paymentV1 "shared/pkg/proto/payment/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func PartToProto(in model.Part) *inventoryV1.Part {
	return &inventoryV1.Part{
		Uuid:          in.Uuid,
		Name:          in.Name,
		Description:   in.Description,
		Price:         in.Price,
		StockQuantity: in.StockQuantity,
		Category:      inventoryV1.Category(in.Category),
		Dimensions:    dimensionsToProto(in.Dimensions),
		Manufacturer:  manufacturerToProto(in.Manufacturer),
		Tags:          append([]string(nil), in.Tags...),
		Metadata:      valuesToProto(in.Metadata),
		CreatedAt:     timeToProto(in.CreatedAt),
		UpdatedAt:     timeToProto(in.UpdatedAt),
	}
}

func PartsToProto(in []model.Part) []*inventoryV1.Part {
	out := make([]*inventoryV1.Part, 0, len(in))
	for _, p := range in {
		out = append(out, PartToProto(p))
	}
	return out
}

func PartsFilterFromProto(in *inventoryV1.PartsFilter) model.PartsFilter {
	if in == nil {
		return model.PartsFilter{}
	}

	categories := make([]model.Category, len(in.Categories))
	for i, c := range in.Categories {
		categories[i] = model.Category(c)
	}

	return model.PartsFilter{
		Uuids:                 append([]string(nil), in.Uuids...),
		Names:                 append([]string(nil), in.Names...),
		Categories:            categories,
		ManufacturerCountries: append([]string(nil), in.ManufacturerCountries...),
		Tags:                  append([]string(nil), in.Tags...),
	}
}

func PartsFromProto(in []*inventoryV1.Part) []model.Part {
	out := make([]model.Part, 0, len(in))
	for _, p := range in{
		out = append(out, PartFromProto(p))
	}
	return out
}

func PartFromProto(in *inventoryV1.Part) model.Part {
	if in == nil {
		return model.Part{}
	}

	return model.Part{
		Uuid:          in.Uuid,
		Name:          in.Name,
		Description:   in.Description,
		Price:         in.Price,
		StockQuantity: in.StockQuantity,
		Category:      model.Category(in.Category),
		Dimensions:    dimensionsFromProto(in.Dimensions),
		Manufacturer:  manufacturerFromProto(in.Manufacturer),
		Tags:          append([]string(nil), in.Tags...),
		Metadata:      valuesFromProto(in.Metadata),
		CreatedAt:     timeFromProto(in.CreatedAt),
		UpdatedAt:     timeFromProto(in.UpdatedAt),
	}
}

func valuesToProto(in map[string]*model.Value) map[string]*inventoryV1.Value {
	if in == nil {
		return nil
	}

	out := make(map[string]*inventoryV1.Value, len(in))
	for k, v := range in {
		if v == nil || v.Kind == nil {
			out[k] = nil
			continue
		}

		switch x := v.Kind.(type) {
		case *model.Value_StringValue:
			out[k] = &inventoryV1.Value{Kind: &inventoryV1.Value_StringValue{StringValue: x.StringValue}}
		case *model.Value_Int64Value:
			out[k] = &inventoryV1.Value{Kind: &inventoryV1.Value_Int64Value{Int64Value: x.Int64Value}}
		case *model.Value_DoubleValue:
			out[k] = &inventoryV1.Value{Kind: &inventoryV1.Value_DoubleValue{DoubleValue: x.DoubleValue}}
		case *model.Value_BoolValue:
			out[k] = &inventoryV1.Value{Kind: &inventoryV1.Value_BoolValue{BoolValue: x.BoolValue}}
		default:
			out[k] = nil
		}
	}
	return out
}

func valuesFromProto(in map[string]*inventoryV1.Value) map[string]*model.Value {
	if in == nil {
		return nil
	}

	out := make(map[string]*model.Value, len(in))
	for k, v := range in {
		if v == nil || v.Kind == nil {
			out[k] = nil
			continue
		}

		switch x := v.Kind.(type) {
		case *inventoryV1.Value_StringValue:
			out[k] = &model.Value{Kind: &model.Value_StringValue{StringValue: x.StringValue}}
		case *inventoryV1.Value_Int64Value:
			out[k] = &model.Value{Kind: &model.Value_Int64Value{Int64Value: x.Int64Value}}
		case *inventoryV1.Value_DoubleValue:
			out[k] = &model.Value{Kind: &model.Value_DoubleValue{DoubleValue: x.DoubleValue}}
		case *inventoryV1.Value_BoolValue:
			out[k] = &model.Value{Kind: &model.Value_BoolValue{BoolValue: x.BoolValue}}
		default:
			out[k] = nil
		}
	}
	return out
}

func dimensionsToProto(in *model.Dimensions) *inventoryV1.Dimensions {
	if in == nil {
		return nil
	}
	return &inventoryV1.Dimensions{
		Length: in.Length,
		Width:  in.Width,
		Height: in.Height,
		Weight: in.Weight,
	}
}

func dimensionsFromProto(in *inventoryV1.Dimensions) *model.Dimensions {
	if in == nil {
		return nil
	}
	return &model.Dimensions{
		Length: in.Length,
		Width:  in.Width,
		Height: in.Height,
		Weight: in.Weight,
	}
}

func manufacturerToProto(in *model.Manufacturer) *inventoryV1.Manufacturer {
	if in == nil {
		return nil
	}
	return &inventoryV1.Manufacturer{
		Name:    in.Name,
		Country: in.Country,
		Website: in.Website,
	}
}

func manufacturerFromProto(in *inventoryV1.Manufacturer) *model.Manufacturer {
	if in == nil {
		return nil
	}
	return &model.Manufacturer{
		Name:    in.Name,
		Country: in.Country,
		Website: in.Website,
	}
}

func timeToProto(in *time.Time) *timestamppb.Timestamp {
	if in == nil {
		return nil
	}
	return timestamppb.New(*in)
}

func timeFromProto(in *timestamppb.Timestamp) *time.Time {
	if in == nil {
		return nil
	}
	t := in.AsTime()
	return &t
}

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

func PaymentMethodModelToProto(in model.PaymentMethod) paymentV1.PaymentMethod {
	switch in {
	case model.PaymentMethod_PAYMENT_METHOD_CARD:
		return paymentV1.PaymentMethod_PAYMENT_METHOD_CARD
	case model.PaymentMethod_PAYMENT_METHOD_SBP:
		return paymentV1.PaymentMethod_PAYMENT_METHOD_SBP
	case model.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD:
		return paymentV1.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD
	case model.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY:
		return paymentV1.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY
	default:
		return paymentV1.PaymentMethod_PAYMENT_METHOD_UNSPECIFIED
	}
}