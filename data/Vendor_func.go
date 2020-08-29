package model

import (
	"VendingMachineTracker/adapter/domain"
	"VendingMachineTracker/domain/usecase/entity"
)

func (v *Vendor) GetID() int32 {
	return int32(v.ID)
}

func (v *Vendor) GetLocation() domain.IAddress {
	ad := entity.NewAddress(v.Country,v.City,v.Line1,v.Line2)
	return ad
}

func (v *Vendor) GetCategory() int16 {
	return int16(v.Category)
}

func (v *Vendor) GetFranchiseID() int32 {
	return int32(v.FranchiseID)
}