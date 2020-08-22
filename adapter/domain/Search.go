package domain

import "time"

type Filter interface{
	GetCountry() string
	GetCity() string
	GetCategory() int16
	GetVendorID() int32
	GetGoodsNameText() string
	GetMinPrice() float64
	GetMaxPrice() float64
	GetStartAt() time.Time
	GetEndAt() time.Time
}
