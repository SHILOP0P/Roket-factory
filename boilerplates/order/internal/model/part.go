package model

import (
	"time"
)

type Part struct {
	Uuid          string                 
	Name          string                 
	Description   string                 
	Price         float64                
	StockQuantity int64                  
	Category      Category               
	Dimensions    *Dimensions            
	Manufacturer  *Manufacturer          
	Tags          []string               
	Metadata      map[string]*Value      
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
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
	Length        float64
	Width         float64
	Height        float64
	Weight        float64
}

type Manufacturer struct {
	Name          string 
	Country       string 
	Website       string 
}

type Value struct {
	Kind isValue_Kind
}

type isValue_Kind interface {
	isValue_Kind()
}

type Value_StringValue struct {
	StringValue string
}
func (*Value_StringValue) isValue_Kind() {}

type Value_Int64Value struct {
	Int64Value int64
}
func (*Value_Int64Value) isValue_Kind() {}

type Value_DoubleValue struct {
	DoubleValue float64
}
func (*Value_DoubleValue) isValue_Kind() {}

type Value_BoolValue struct {
	BoolValue bool
}
func (*Value_BoolValue) isValue_Kind() {}

type PartsFilter struct {
	Uuids                 []string   
	Names                 []string   
	Categories            []Category 
	ManufacturerCountries []string  
	Tags                  []string   
}
