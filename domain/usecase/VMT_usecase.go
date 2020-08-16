package usecase

import (
	"VendingMachineTracker/adapter/domain"
	"VendingMachineTracker/domain/repository"
	"errors"
)

type VmtUseCase struct {
	repo repository.VmtRepository
}

func (v *VmtUseCase) InitRepository() {
	v.repo.InitRepository()
}


//Get Record From Outer, validate and save it to repository
func (v *VmtUseCase) SaveSalesRecord(record domain.ISalesRecords) error {
	if err := v.repo.SaveSalesRecord(record);err!= nil {
		return errors.New("repository Error")
	}
	return nil
}

func (v *VmtUseCase) ViewRecordData(query interface{}) interface{} {
	return nil
}
