package part

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"inventory/internal/model"
)

func (s *ServiceSuite) TestGetSuccess(){
	var(
		uuid = gofakeit.UUID()
		name = gofakeit.Name()
		description = gofakeit.Paragraph(3, 5, 5, " ")
		price = gofakeit.Float64()
		stockQuantity = gofakeit.Int64()
		category = int32(gofakeit.Number(0, 4))
		length = gofakeit.Float64()
		width = gofakeit.Float64()
		height = gofakeit.Float64()
		weight = gofakeit.Float64()
		nameManufacture = gofakeit.Name()
		country = gofakeit.Country()
		website = gofakeit.Word()
		n = gofakeit.Number(0, 10)
		tags = make([]string, n)
		metadata = fakeMetadata()
		createdAt = time.Now()
		updatedAt = time.Now()

		part = model.Part{
			Uuid: uuid,
			Name: name,
			Description: description,
			Price: price,
			StockQuantity: stockQuantity,
			Category: model.Category(category),
			Dimensions: &model.Dimensions{
				Length: length,
				Width: width,
				Weight: weight,
				Height: height,
			},
			Manufacturer: &model.Manufacturer{
				Name: nameManufacture,
				Country: country,
				Website: website,
			},
			Tags: tags,
			Metadata: metadata,
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
		}

	)
	for i := 0; i < n; i++ {
			part.Tags[i] = gofakeit.Word()
	}

	s.inventoryRepository.On("GetPart", s.ctx, uuid).Return(part, nil)
	
	res, err := s.service.GetPart(s.ctx, uuid)
	s.NoError(err)
	s.Equal(part, res)
}

func (s *ServiceSuite) TestGetError(){
	var(
		repoError = gofakeit.Error()
		uuid = gofakeit.UUID()
	)
	s.inventoryRepository.On("GetPart", s.ctx, uuid).Return(model.Part{}, repoError)
	
	res, err := s.service.GetPart(s.ctx, uuid)
	s.Error(err)
	s.ErrorIs(err, repoError)
	s.Empty(res)
}




func fakeMetadata() map[string]*model.Value{
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