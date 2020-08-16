package entity

import "VendingMachineTracker/adapter/domain"

type Vendor struct {
	ID       int32
	Location domain.IAddress
	Category int16
}

func (v *Vendor) GetID() int32 {
	return v.ID
}

func (v *Vendor) GetLocation() domain.IAddress {
	return v.Location
}

func (v *Vendor) GetCategory() int16 {
	return v.Category
}

type Address struct {
	Country string
	City    string
	Line1   string
	Line2   string
}

func (a *Address) GetCountry() string {
	return a.Country
}

func (a *Address) GetCity() string {
	return a.City
}

func (a *Address) GetLine1() string {
	return a.Line1
}

func (a *Address) GetLine2() string {
	return a.Line2
}

