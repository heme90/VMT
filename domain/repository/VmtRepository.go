package repository

import (
	"VendingMachineTracker/adapter/domain"
	model "VendingMachineTracker/data"
	"VendingMachineTracker/domain/usecase/entity"
	"errors"
	"github.com/go-pg/pg"
	"sync"
	"time"
)

var once sync.Once

type VmtRepository struct {
	conn *pg.DB
}

func (vr * VmtRepository) InitRepository() {
	once.Do(func(){
		vr.conn = pg.Connect(&pg.Options{
			Addr:     ":5432",
			User:     "user",
			Password: "pass",
			Database: "db_name",
		})
	})

}

func (vr * VmtRepository) GetInstance() *pg.DB {
	if vr.conn == nil {
		panic("Initialize Repository Connection First!")
	}

	return vr.conn
}

func (vr *VmtRepository) SaveSalesRecord(record domain.ISalesRecords) error {
	tx,err := vr.conn.Begin()
	if err != nil {
		return errors.New("There's Error In Begin TRANSACTION!")
	}
	//TODO implement Model
	if _,err := tx.Model(model.SaleRecord{}).Insert(record); err!= nil {
		err := tx.Rollback()
		if err != nil {
			return errors.New("error in rollback transaction!")
		}
		return errors.New("error in save record : ")
	}

	return nil
}
//TODO error 패키지로 정의?
func (vr *VmtRepository) GetSalesRecords(filter domain.Filter) (domain.SearchResponse,error) {
	response := entity.NewSearchResponse()
	q :=vr.conn.Model(model.SaleRecord{})
	//TODO pg.Indent.Model ..... Db 구성, 모델 생성 후 쿼리문 수정
	//TODO 조인쿼리 orm function 으로 분리
	if filter.GetCountry() != "" {
		q = q.Where("? = ?")
	}
	if filter.GetCity() != "" {
		q = q.Where("? = ?")
	}
	if filter.GetCategory() != 0 {
		q = q.Where("? = ?")
	}
	if filter.GetVendorID() != 0 {
		q = q.Where("? = ?")
	}
	if filter.GetGoodsNameText() != "" {
		q = q.Where("? = ?")
	}
	if filter.GetMinPrice() != 0 {
		q = q.Where("? = ?")
	}
	if filter.GetMaxPrice() != 0 {
		q = q.Where("? = ?")
	}
	if filter.GetStartAt() != time.Date(1,1,1, 0,0,0 ,0,nil)  {
		q = q.Where("? = ?")
	}
	if filter.GetEndAt() != time.Date(1,1,1, 0,0,0 ,0,nil)  {
		q = q.Where("? = ?")
	}
	if err := q.Select(&response);err!=nil {
		return nil,err
	}

	return response, nil
}