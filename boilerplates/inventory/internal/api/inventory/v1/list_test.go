package v1

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"inventory/internal/converter"
	"inventory/internal/model"
	inventory "shared/pkg/proto/inventory/v1"
)

func (s *APISuite) TestListPartsSuccess() {
    listPart := fakeListParts()
    req := &inventory.ListPartsRequest{}


	expectedPart := converter.PartsToProto(listPart)

	s.inventoryService.On("ListParts", s.ctx, converter.PartsFilterFromProto(req.GetFilter())).Return(listPart, nil)

	res, err := s.api.ListParts(s.ctx, req)
	s.NoError(err)
	s.Require().Len(res.GetParts(), len(expectedPart))
    for i := range expectedPart   {
        resPart := res.GetParts()[i]
        expPart := expectedPart[i]

        s.Require().Equal(expPart.GetUuid(), resPart.GetUuid())
        s.Require().Equal(expPart.GetName(), resPart.GetName())
        s.Require().Equal(expPart.GetDescription(), resPart.GetDescription())
        s.Require().Equal(expPart.GetPrice(), resPart.GetPrice())
        s.Require().Equal(expPart.GetStockQuantity(), resPart.GetStockQuantity())
        s.Require().Equal(expPart.GetCategory(), resPart.GetCategory())
        s.Require().Equal(expPart.GetDimensions(), resPart.GetDimensions())
        s.Require().Equal(expPart.GetManufacturer(), resPart.GetManufacturer())
        s.Require().Equal(expPart.GetTags(), resPart.GetTags())
    }

}

func (s *APISuite) TestListPartsError() {
	var (
        serviceError = gofakeit.Error()
        filter = model.PartsFilter{}
		req = &inventory.ListPartsRequest{}
	)
	s.inventoryService.On("ListParts", s.ctx, filter).Return(nil, serviceError)

	res, err := s.api.ListParts(s.ctx, req)
	s.Error(err)
	s.ErrorIs(err, serviceError)
	s.Empty(res)
}


func fakeListParts() []model.Part{
    n := gofakeit.Number(1, 6)

    listParts := make([]model.Part, n)

    for i := 0; i < n; i++ {
        uuid            := gofakeit.UUID()
		name            := gofakeit.Name()
		description     := gofakeit.Paragraph(3, 5, 5, " ")
		price           := gofakeit.Float64()
		stockQuantity   := gofakeit.Int64()
		category        := int32(gofakeit.Number(0, 4))
		length          := gofakeit.Float64()
		width           := gofakeit.Float64()
		height          := gofakeit.Float64()
		weight          := gofakeit.Float64()
		nameManufacture := gofakeit.Name()
		country         := gofakeit.Country()
		website         := gofakeit.Word()
		n               := gofakeit.Number(0, 10)
		tags            := make([]string, n)
		metadata        := fakeMetadata()
		createdAt       := time.Now()
		updatedAt       := time.Now()

        modelPart := model.Part{
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
        listParts[i] = modelPart 
    }
    return listParts
}