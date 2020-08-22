package entity

import (
	"VendingMachineTracker/adapter/domain"
	"time"
)

//1. 상품명
//2. 가격
//3. 상품 카테고리
//4. 판매 시각
//5. 지역

type SalesRecord struct {
	VendorID  int32
	GoodsName string
	Price     float64
	Currency  string
	CreatedAt time.Time
}

func NewSalesRecord(vendorID int32, goodsName string, price float64, currency string, createdAt time.Time) *SalesRecord {
	return &SalesRecord{VendorID: vendorID, GoodsName: goodsName, Price: price, Currency: currency, CreatedAt: createdAt}
}

func (s *SalesRecord) GetVendorID() int32 {
	return s.VendorID
}

func (s *SalesRecord) GetGoodsName() string {
	return s.GoodsName
}

func (s *SalesRecord) GetPrice() float64 {
	return s.Price
}

func (s *SalesRecord) GetCurrency() string {
	return s.Currency
}

func (s *SalesRecord) GetCreatedAt() time.Time {
	return s.CreatedAt
}

type SearchResponse struct {
	records []SalesRecord
}

func NewSearchResponse() *SearchResponse {
	return &SearchResponse{}
}

func (s *SearchResponse) GetSearchRecords() []domain.ISalesRecords {
	res := make([]domain.ISalesRecords,0)
	for _,v := range s.records {
		res = append(res,&v)
	}
	return res
}
