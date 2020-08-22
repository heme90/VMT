package entity

import "time"

type SearchFilter struct {
	country  string
	city     string
	category int16
	vendorID int32
	text     string
	minPrice float64
	maxPrice float64
	startAt  time.Time
	endAt    time.Time
}

func NewSearchFilter(country string, city string, category int16, vendorID int32, nameText string, minPrice float64, maxPrice float64, startAt time.Time, endAt time.Time) *SearchFilter {
	return &SearchFilter{country: country, city: city, category: category, vendorID: vendorID, text: nameText, minPrice: minPrice, maxPrice: maxPrice, startAt: startAt, endAt: endAt}
}

func (s *SearchFilter) GetCountry() string {
	return s.country
}

func (s *SearchFilter) GetCity() string {
	return s.city
}

func (s *SearchFilter) GetCategory() int16{
	return s.category
}

func (s *SearchFilter) GetVendorID() int32 {
	return s.vendorID
}

func (s *SearchFilter) GetGoodsNameText() string {
	return s.text
}

func (s *SearchFilter) GetMinPrice() float64 {
	return s.minPrice
}

func (s *SearchFilter) GetMaxPrice() float64 {
	return s.maxPrice
}

func (s *SearchFilter) GetStartAt() time.Time {
	return s.startAt
}

func (s *SearchFilter) GetEndAt() time.Time {
	return s.endAt
}

