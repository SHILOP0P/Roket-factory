package v1

import (
	"context"
	"order/internal/model"
	"order/internal/client/converter"
	inventory_v1 "shared/pkg/proto/inventory/v1"

)

func (c *inventoryClient) GetInventoryModels(ctx context.Context, uuids []string)([]model.Part, error){
	req := &inventory_v1.ListPartsRequest{
		Filter: &inventory_v1.PartsFilter{
			Uuids: uuids,
		},
	}
	res, err := c.client.ListParts(ctx, req)
	if err!=nil{
		return nil, err
	}
	parts := converter.PartsFromProto(res.Parts)
	return parts, nil
}