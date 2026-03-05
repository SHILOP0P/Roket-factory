package part

import (
	repoModel "inventory/internal/repository/model"
)

func seedParts() map[string]*repoModel.Part {
	return map[string]*repoModel.Part{
		"11111111-1111-1111-1111-111111111111": {
			Uuid:          "11111111-1111-1111-1111-111111111111",
			Name:          "Main Engine X1",
			Description:   "Основной двигатель корабля",
			Price:         120000.50,
			StockQuantity: 10,
			Category:      repoModel.Category_CATEGORY_ENGINE,
			Dimensions: &repoModel.Dimensions{
				Length: 320.0,
				Width:  180.0,
				Height: 190.0,
				Weight: 950.0,
			},
			Manufacturer: &repoModel.Manufacturer{
				Name:    "Space Motors GmbH",
				Country: "Germany",
				Website: "https://example.com/spacemotors",
			},
			Tags: []string{"engine", "main", "heavy"},
			Metadata: map[string]*repoModel.Value{
				"series": {Kind: &repoModel.Value_StringValue{StringValue: "X"}},
			},
		},
		"22222222-2222-2222-2222-222222222222": {
			Uuid:          "22222222-2222-2222-2222-222222222222",
			Name:          "Fuel Tank F2",
			Description:   "Топливный бак",
			Price:         45000.00,
			StockQuantity: 25,
			Category:      repoModel.Category_CATEGORY_FUEL,
			Dimensions: &repoModel.Dimensions{
				Length: 210.0,
				Width:  160.0,
				Height: 240.0,
				Weight: 400.0,
			},
			Manufacturer: &repoModel.Manufacturer{
				Name:    "Orbital Fuel Inc.",
				Country: "USA",
				Website: "https://example.com/orbitalfuel",
			},
			Tags: []string{"fuel", "tank"},
			Metadata: map[string]*repoModel.Value{
				"pressure_rating": {Kind: &repoModel.Value_Int64Value{Int64Value: 300}},
			},
		},
	}
}