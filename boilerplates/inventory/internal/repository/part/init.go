package part

import (
	"time"

	repoModel "inventory/internal/repository/model"
)

func seedParts() []repoModel.Part {
	now := time.Now()

	return []repoModel.Part{
		{
			Uuid:          "11111111-1111-1111-1111-111111111111",
			Name:          "Merlin 1D Engine",
			Description:   "Liquid rocket engine used on the Falcon 9 first stage.",
			Price:         1000000,
			StockQuantity: 4,
			Category:      repoModel.Category_CATEGORY_ENGINE,
			Dimensions: &repoModel.Dimensions{
				Length: 4.0,
				Width:  1.2,
				Height: 1.2,
				Weight: 470,
			},
			Manufacturer: &repoModel.Manufacturer{
				Name:    "SpaceX",
				Country: "USA",
				Website: "https://www.spacex.com",
			},
			Tags: []string{"engine", "rocket", "liquid"},
			Metadata: map[string]*repoModel.Value{
				"thrust_kN": {
					DoubleValue: float64Ptr(845),
				},
				"fuel": {
					StringValue: stringPtr("RP-1 / LOX"),
				},
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		{
			Uuid:          "22222222-2222-2222-2222-222222222222",
			Name:          "Raptor Engine",
			Description:   "Methane-fueled full-flow staged combustion rocket engine.",
			Price:         2000000,
			StockQuantity: 2,
			Category:      repoModel.Category_CATEGORY_ENGINE,
			Dimensions: &repoModel.Dimensions{
				Length: 3.1,
				Width:  1.3,
				Height: 1.3,
				Weight: 1600,
			},
			Manufacturer: &repoModel.Manufacturer{
				Name:    "SpaceX",
				Country: "USA",
				Website: "https://www.spacex.com",
			},
			Tags: []string{"engine", "methane", "starship"},
			Metadata: map[string]*repoModel.Value{
				"thrust_kN": {
					DoubleValue: float64Ptr(2300),
				},
				"fuel": {
					StringValue: stringPtr("CH4 / LOX"),
				},
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		{
			Uuid:          "33333333-3333-3333-3333-333333333333",
			Name:          "Soyuz Fuel Tank",
			Description:   "Propellant tank section for medium-lift launch vehicle.",
			Price:         350000,
			StockQuantity: 6,
			Category:      repoModel.Category_CATEGORY_FUEL,
			Dimensions: &repoModel.Dimensions{
				Length: 6.5,
				Width:  2.7,
				Height: 2.7,
				Weight: 1200,
			},
			Manufacturer: &repoModel.Manufacturer{
				Name:    "Progress Rocket Space Centre",
				Country: "Russia",
				Website: "https://www.samspace.ru",
			},
			Tags: []string{"fuel", "tank", "soyuz"},
			Metadata: map[string]*repoModel.Value{
				"material": {
					StringValue: stringPtr("Aluminum alloy"),
				},
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		{
			Uuid:          "44444444-4444-4444-4444-444444444444",
			Name:          "Dragon Porthole",
			Description:   "Spacecraft-grade window assembly for crew capsule.",
			Price:         120000,
			StockQuantity: 8,
			Category:      repoModel.Category_CATEGORY_PORTHOLE,
			Dimensions: &repoModel.Dimensions{
				Length: 0.7,
				Width:  0.7,
				Height: 0.2,
				Weight: 45,
			},
			Manufacturer: &repoModel.Manufacturer{
				Name:    "SpaceX",
				Country: "USA",
				Website: "https://www.spacex.com",
			},
			Tags: []string{"porthole", "window", "capsule"},
			Metadata: map[string]*repoModel.Value{
				"radiation_resistant": {
					BoolValue: boolPtr(true),
				},
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		{
			Uuid:          "55555555-5555-5555-5555-555555555555",
			Name:          "Buran Wing Panel",
			Description:   "Reusable orbiter wing thermal-protected structural panel.",
			Price:         500000,
			StockQuantity: 1,
			Category:      repoModel.Category_CATEGORY_WING,
			Dimensions: &repoModel.Dimensions{
				Length: 5.8,
				Width:  2.1,
				Height: 0.4,
				Weight: 780,
			},
			Manufacturer: &repoModel.Manufacturer{
				Name:    "NPO Molniya",
				Country: "Russia",
				Website: "https://www.buran.ru",
			},
			Tags: []string{"wing", "orbiter", "thermal"},
			Metadata: map[string]*repoModel.Value{
				"reusable": {
					BoolValue: boolPtr(true),
				},
			},
			CreatedAt: &now,
			UpdatedAt: &now,
		},
	}
}

func stringPtr(v string) *string {
	return &v
}

func float64Ptr(v float64) *float64 {
	return &v
}

func boolPtr(v bool) *bool {
	return &v
}
