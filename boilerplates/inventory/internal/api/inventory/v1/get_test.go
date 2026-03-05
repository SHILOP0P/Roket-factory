package v1

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"inventory/internal/converter"
	"inventory/internal/model"
	inventory "shared/pkg/proto/inventory/v1"
)

func (s *APISuite) TestGetSuccess() {
	var (
		uuid            = gofakeit.UUID()
		name            = gofakeit.Name()
		description     = gofakeit.Paragraph(3, 5, 5, " ")
		price           = gofakeit.Float64()
		stockQuantity   = gofakeit.Int64()
		category        = int32(gofakeit.Number(0, 4))
		length          = gofakeit.Float64()
		width           = gofakeit.Float64()
		height          = gofakeit.Float64()
		weight          = gofakeit.Float64()
		nameManufacture = gofakeit.Name()
		country         = gofakeit.Country()
		website         = gofakeit.Word()
		n               = gofakeit.Number(0, 10)
		tags            = make([]string, n)
		metadata        = fakeMetadata()
		createdAt       = time.Now()
		updatedAt       = time.Now()

		req = &inventory.GetPartRequest{
			Uuid: uuid,
		}

		modelPart = model.Part{
			Uuid:          uuid,
			Name:          name,
			Description:   description,
			Price:         price,
			StockQuantity: stockQuantity,
			Category:      model.Category(category),
			Dimensions: &model.Dimensions{
				Length: length,
				Width:  width,
				Weight: weight,
				Height: height,
			},
			Manufacturer: &model.Manufacturer{
				Name:    nameManufacture,
				Country: country,
				Website: website,
			},
			Tags:      tags,
			Metadata:  metadata,
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
		}
	)
	for i := 0; i < n; i++ {
		modelPart.Tags[i] = gofakeit.Word()
	}

	expectedPart := converter.PartToProto(modelPart)

	s.inventoryService.On("GetPart", s.ctx, uuid).Return(modelPart, nil)

	res, err := s.api.GetPart(s.ctx, req)
	s.NoError(err)
	s.Require().NotNil(res)
	s.Require().Equal(expectedPart.Uuid, res.GetPart().GetUuid())
	s.Require().Equal(expectedPart.Name, res.GetPart().GetName())
	s.Require().Equal(expectedPart.Description, res.GetPart().GetDescription())
	s.Require().Equal(expectedPart.Price, res.GetPart().GetPrice())
	s.Require().Equal(expectedPart.StockQuantity, res.GetPart().GetStockQuantity())
	s.Require().Equal(expectedPart.Category, res.GetPart().GetCategory())
	s.Require().Equal(expectedPart.Dimensions, res.GetPart().GetDimensions())
	s.Require().Equal(expectedPart.Manufacturer, res.GetPart().GetManufacturer())
	s.Require().Equal(expectedPart.Tags, res.GetPart().GetTags())
}

func (s *APISuite) TestGetError() {
	var (
		uuid = gofakeit.UUID()

		req = &inventory.GetPartRequest{
			Uuid: uuid,
		}
	)
	s.inventoryService.On("GetPart", s.ctx, uuid).Return(model.Part{}, model.ErrPartNotFound)

	res, err := s.api.GetPart(s.ctx, req)
	s.Require().Error(err)
	s.Require().Nil(res)

	st, ok := status.FromError(err)
	s.Require().True(ok)
	s.Require().Equal(codes.NotFound, st.Code())
}

func fakeMetadata() map[string]*model.Value {
	n := gofakeit.Number(0, 4)
	m := make(map[string]*model.Value, n)

	for i := 0; i < n; i++ {
		key := gofakeit.Word()
		switch gofakeit.Number(0, 3) {
		case 0:
			m[key] = &model.Value{Kind: &model.Value_StringValue{StringValue: gofakeit.Word()}}
		case 1:
			m[key] = &model.Value{Kind: &model.Value_Int64Value{Int64Value: gofakeit.Int64()}}
		case 2:
			m[key] = &model.Value{Kind: &model.Value_DoubleValue{DoubleValue: gofakeit.Float64Range(0, 1000)}}
		default:
			m[key] = &model.Value{Kind: &model.Value_BoolValue{BoolValue: gofakeit.Bool()}}
		}
	}

	return m
}
