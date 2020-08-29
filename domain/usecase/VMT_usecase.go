package usecase

import (
	dto "VendingMachineTracker/adapter/domain"
	"VendingMachineTracker/domain/repository"
	"errors"
	"log"
)

type VmtUseCase struct {
	repo repository.VmtRepository
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



	response,err := v.repo.GetSalesRecords(searchFilter)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return response
}

func (v *VmtUseCase) AuthFranchiseAdmin(apikey string) error {
	panic("implement me")
}
