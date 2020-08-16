package domain

import "time"

//type SalesRecord struct {
//	VendorID int32
//	GoodsName string
//	Price float64
//	Currency string
//	CreatedAt time.Time
//}

type ISalesRecords interface {
	GetVendorID() int32
	GetGoodsName() string
	GetPrice() float64
	GetCurrency() string
	GetCreatedAt() time.Time
}