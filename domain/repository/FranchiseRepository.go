package repository

import (
	model "VendingMachineTracker/data"
)

func (vr * VmtRepository) GetFranchiseAuth(apiKey string) int32 {
	res := model.Franchise{}
	if err:= vr.conn.Model(model.Franchise{}).
		Where("?=?",model.Columns.Franchise.Apikey,apiKey).
		Select(&res);err!=nil{
		return 0
	}
	return res.GetID()
}