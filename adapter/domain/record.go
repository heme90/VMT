package domain

import "time"

type ISalesRecords interface {
	GetVendorID() int32
	GetGoodsName() string
	GetPrice() float64
	GetCurrency() string
	GetCreatedAt() time.Time
}

type SearchResponse interface{
	GetSearchRecords() []ISalesRecords
}