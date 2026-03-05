package part

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"inventory/internal/model"
)

func (s *ServiceSuite) TestListPartsSuccess(){
	filter := model.PartsFilter{}
	expected := fakeParts(gofakeit.Number(1, 5))

	s.inventoryRepository.On("ListParts", s.ctx, filter).Return(expected, nil)

	res, err := s.service.ListParts(s.ctx, filter)
	s.NoError(err)
	s.Equal(expected, res)
}

func (s *ServiceSuite) TestListPartsError(){
	var(
		repoError = gofakeit.Error()
		filter = model.PartsFilter{}
	)
	s.inventoryRepository.On("ListParts", s.ctx, filter).Return(nil, repoError)
	
	res, err := s.service.ListParts(s.ctx, filter)
	s.Error(err)
	s.ErrorIs(err, repoError)
	s.Empty(res)
}

func fakeParts(n int) []model.Part {
	parts := make([]model.Part, 0, n)

	for i := 0; i < n; i++ {
		createdAt := time.Now()
		updatedAt := time.Now()

		part := model.Part{
			Uuid:          gofakeit.UUID(),
			Name:          gofakeit.Name(),
			Description:   gofakeit.Sentence(10),
			Price:         gofakeit.Float64Range(1, 10000),
			StockQuantity: gofakeit.Int64(),
			Category:      model.Category(gofakeit.Number(0, 4)),
			Dimensions: &model.Dimensions{
				Length: gofakeit.Float64Range(1, 100),
				Width:  gofakeit.Float64Range(1, 100),
				Height: gofakeit.Float64Range(1, 100),
				Weight: gofakeit.Float64Range(1, 100),
			},
			Manufacturer: &model.Manufacturer{
				Name:    gofakeit.Company(),
				Country: gofakeit.Country(),
				Website: gofakeit.URL(),
			},
			Tags:      []string{gofakeit.Word(), gofakeit.Word()},
			Metadata:  fakeMetadata(),
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
		}

		parts = append(parts, part)
	}

	return parts
}
