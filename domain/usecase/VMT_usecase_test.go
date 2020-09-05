package usecase

import (
	"VendingMachineTracker/adapter/domain"
	repository2 "VendingMachineTracker/domain/repository"
	"VendingMachineTracker/domain/usecase/entity"
	"fmt"
	"testing"
	"time"
)

var reposit domain.Repository

func init() {


}


//TODO json 파일 만들어서 mock 테스트 작성
func TestVmtUseCase_SaveSalesRecord(t *testing.T) {
	repo := repository2.VmtRepository{}
	reposit = repo.InitRepository()
	rec := entity.NewSalesRecord(1,"pocari",700,"KRW",time.Now())
	use := NewVmtUseCase(reposit)
	if err := use.SaveSalesRecord(rec);err != nil {
		fmt.Errorf("error %v",err)
	}
}