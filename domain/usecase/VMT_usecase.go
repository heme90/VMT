package usecase

import (
	dto "VendingMachineTracker/adapter/domain"
	"errors"
	"log"
)

type VmtUseCase struct {
	repo dto.Repository
}





func (v *VmtUseCase) InitRepository() {
	v.repo.InitRepository()
}


//Get Record From Outer, validate and save it to repository
func (v *VmtUseCase) SaveSalesRecord(record dto.ISalesRecords) error {
	if err := v.repo.SaveSalesRecord(record);err!= nil {
		return errors.New("repository Error")
	}
	return nil
}

func (v *VmtUseCase) ViewRecordData(searchFilter dto.Filter) dto.SearchResponse {
	response,err := v.repo.SearchSalesRecords(searchFilter)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return response
}

func (v *VmtUseCase) AuthFranchiseAdmin(apikey string) (int32,error) {
	var franchiseID int32
	if franchiseID = v.repo.GetFranchiseAuth(apikey);franchiseID==0{
		return 0, errors.New("auth fail")
	}
	return franchiseID,nil
}
