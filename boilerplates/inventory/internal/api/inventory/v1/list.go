package v1

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"inventory/internal/model"
	"inventory/internal/converter"
	inventoryV1 "shared/pkg/proto/inventory/v1"
)

func (a *api) ListParts(ctx context.Context, req *inventoryV1.ListPartsRequest)(*inventoryV1.ListPartsResponse, error) {
	parts, err := a.inventoryService.ListParts(ctx, converter.PartsFilterFromProto(req.GetFilter()))
	if err!=nil{
		if errors.Is(err, model.ErrPartNotFound) {
			return nil, status.Errorf(codes.OK, "Parts not found")
		}
		return nil, err
	}

	return &inventoryV1.ListPartsResponse{
		Parts: converter.PartsToProto(parts),
	}, nil
}