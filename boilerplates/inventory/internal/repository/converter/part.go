package converter

import (
	repoModel "inventory/internal/repository/model"
	"inventory/internal/model"
)

func PartToRepoModel(part model.Part) repoModel.Part{
	return repoModel.Part{
		Uuid: part.Uuid,
		Name: part.Name,
		Description: part.Description,
		Price: part.Price,
		StockQuantity: part.StockQuantity,
		Category: repoModel.Category(part.Category),
		Dimensions: (*repoModel.Dimensions)(part.Dimensions),
		Manufacturer: (*repoModel.Manufacturer)(part.Manufacturer),
		Tags: part.Tags,
		Metadata: ValuesToRepo(part.Metadata),
		CreatedAt: part.CreatedAt,
		UpdatedAt: part.UpdatedAt,
	}
}

func PartToModel(part *repoModel.Part)model.Part{
	return model.Part{
		Uuid: part.Uuid,
		Name: part.Name,
		Description: part.Description,
		Price: part.Price,
		StockQuantity: part.StockQuantity,
		Category: model.Category(part.Category),
		Dimensions: (*model.Dimensions)(part.Dimensions),
		Manufacturer: (*model.Manufacturer)(part.Manufacturer),
		Tags: part.Tags,
		Metadata: ValuesToModel(part.Metadata),
		CreatedAt: part.CreatedAt,
		UpdatedAt: part.UpdatedAt,
	}
	
}

func PartsToModel(parts []repoModel.Part) []model.Part{
	out := make([]model.Part, 0, len(parts))
	for _, part:=range parts{
		out = append(out, PartToModel(&part))
	}
	return out
}


func ValuesToRepo(in map[string]*model.Value) map[string]*repoModel.Value {
	if in == nil {
		return nil
	}
	out := make(map[string]*repoModel.Value, len(in))
	for k, v := range in {
		if v == nil {
			out[k] = nil
			continue
		}
		out[k] = ValueToRepo(in[k])
	}
	return out
}

func ValuesToModel(in map[string]*repoModel.Value) map[string]*model.Value {
	if in == nil {
		return nil
	}
	out := make(map[string]*model.Value, len(in))
	for k, v := range in {
		if v == nil {
			out[k] = nil
			continue
		}
		out[k] = ValueToModel(in[k])
	}
	return out
}

func PartsFilterToRepoModel(f model.PartsFilter) repoModel.PartsFilter {
	categories := make([]repoModel.Category, len(f.Categories))
	for i, c := range f.Categories {
		categories[i] = repoModel.Category(c)
	}

	return repoModel.PartsFilter{
		Uuids:                 append([]string(nil), f.Uuids...),
		Names:                 append([]string(nil), f.Names...),
		Categories:            categories,
		ManufacturerCountries: append([]string(nil), f.ManufacturerCountries...),
		Tags:                  append([]string(nil), f.Tags...),
	}
}

func PartsFilterToModel(f repoModel.PartsFilter) model.PartsFilter {
	categories := make([]model.Category, len(f.Categories))
	for i, c := range f.Categories {
		categories[i] = model.Category(c)
	}

	return model.PartsFilter{
		Uuids:                 append([]string(nil), f.Uuids...),
		Names:                 append([]string(nil), f.Names...),
		Categories:            categories,
		ManufacturerCountries: append([]string(nil), f.ManufacturerCountries...),
		Tags:                  append([]string(nil), f.Tags...),
	}
}


func ValueToRepo(v *model.Value) *repoModel.Value {
	if v == nil || v.Kind == nil {
		return nil
	}
	switch x := v.Kind.(type) {
	case *model.Value_StringValue:
		return &repoModel.Value{StringValue: &x.StringValue}
	case *model.Value_Int64Value:
		return &repoModel.Value{Int64Value: &x.Int64Value}
	case *model.Value_DoubleValue:
		return &repoModel.Value{DoubleValue: &x.DoubleValue}
	case *model.Value_BoolValue:
		return &repoModel.Value{BoolValue: &x.BoolValue}
	default:
		return nil
	}
}

func ValueToModel(v *repoModel.Value) *model.Value {
	if v == nil {
		return nil
	}
	if v.StringValue!=nil{
		return &model.Value{Kind: &model.Value_StringValue{StringValue: *v.StringValue}}
	}
	if v.Int64Value!=nil{
		return &model.Value{Kind: &model.Value_Int64Value{Int64Value: *v.Int64Value}}
	}
	if v.DoubleValue!=nil{
		return &model.Value{Kind: &model.Value_DoubleValue{DoubleValue: *v.DoubleValue}}
	}
	if v.BoolValue!=nil{
		return &model.Value{Kind: &model.Value_BoolValue{BoolValue: *v.BoolValue}}
	}
	return nil
}
