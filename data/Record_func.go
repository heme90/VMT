package model

import "time"



func (s *SaleRecord) GetID() int64 {
	return int64(s.ID)
}

func (s *SaleRecord) GetVendorID() int32 {
	return int32(s.Vendorid)
}

func (s *SaleRecord) GetGoodsName() string {
	return s.Goodsname
}

func (s *SaleRecord) GetPrice() float64 {
	return s.Price
}

func (s *SaleRecord) GetCurrency() string {
	return s.Currency
}

func (s SaleRecord) GetCreatedAt() time.Time {
	return s.Createdat
}