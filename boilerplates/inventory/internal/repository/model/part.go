package model

import (
	"time"
)

type Part struct {
	Uuid          string `bson:"uuid,omitempty"`
	Name          string `bson:"name"`
	Description   string `bson:"description"`
	Price         float64 `bson:"price"`
	StockQuantity int64 `bson:"stock_quantity"`
	Category      Category `bson:"category"`
	Dimensions    *Dimensions `bson:"dimensions"`
	Manufacturer  *Manufacturer `bson:"manufacturer"`
	Tags          []string `bson:"tags"`
	Metadata      map[string]*Value `bson:"metadata"`
	CreatedAt     *time.Time `bson:"created_at"`
	UpdatedAt     *time.Time `bson:"updated_at"`
}

type Category int32

const (
	Category_CATEGORY_UNSPECIFIED Category = 0
	Category_CATEGORY_ENGINE      Category = 1
	Category_CATEGORY_FUEL        Category = 2
	Category_CATEGORY_PORTHOLE    Category = 3
	Category_CATEGORY_WING        Category = 4
)

type Dimensions struct {
	Length        float64 `bson:"length"`
	Width         float64 `bson:"width"`
	Height        float64 `bson:"height"`
	Weight        float64 `bson:"weight"`
}

type Manufacturer struct {
	Name          string `bson:"name"`
	Country       string `bson:"country"`
	Website       string `bson:"website"`
}

type Value struct {
	StringValue *string `bson:"string_value"`
	Int64Value *int64 `bson:"int64_value"`
	DoubleValue *float64 `bson:"double_value"`
	BoolValue *bool `bson:"bool_value"`
}


type PartsFilter struct {
	Uuids                 []string
	Names                 []string
	Categories            []Category
	ManufacturerCountries []string
	Tags                  []string
}

