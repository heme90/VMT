package repository

import (
	"VendingMachineTracker/adapter/domain"
	model "VendingMachineTracker/data"
	"VendingMachineTracker/domain/usecase/entity"
	"errors"
	"github.com/go-pg/pg"
	"log"
	"sync"
	"time"
)

var once sync.Once

type VmtRepository struct {
	conn *pg.DB
}

func (vr *VmtRepository) InitRepository() domain.Repository {
	once.Do(func(){
		//TODO usecase 작성 이후 상수처리하여 env.json 만들기
		vr.conn = pg.Connect(&pg.Options{
			Addr:     ":5432",
			User:     "vmtUser",
			Password: "07140714",
			Database: "vmt_dev",
		})
	})
	return vr
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
	if err := tx.Commit();err!=nil {
		log.Fatal(err)
		return err
	}
	return nil
}
//TODO error 패키지로 정의?
func (vr *VmtRepository) SearchSalesRecords(filter domain.Filter) (domain.SearchResponse,error) {
	response := entity.NewSearchResponse()
	q :=vr.conn.Model(model.SaleRecord{})
	//TODO pg.Indent.Model ..... Db 구성, 모델 생성 후 쿼리문 수정
	//TODO 조인쿼리 orm function 으로 분리

	if filter.GetVendorID() != 0 {
		q = q.Where("? = ?",model.Columns.SaleRecord.Vendorid,filter.GetVendorID())
	}
	if filter.GetGoodsNameText() != "" {
		q = q.Where("? = ?",model.Columns.SaleRecord.Goodsname,filter.GetGoodsNameText())
	}
	if filter.GetMinPrice() != 0 {
		q = q.Where("? > ?",model.Columns.SaleRecord.Price,filter.GetMinPrice())
	}
	if filter.GetMaxPrice() != 0 {
		q = q.Where("? < ?",model.Columns.SaleRecord.Price,filter.GetMaxPrice())
	}
	if filter.GetStartAt() != time.Date(1,1,1, 0,0,0 ,0,nil)  {
		q = q.Where("? > ?",model.Columns.SaleRecord.Createdat,filter.GetStartAt())
	}
	if filter.GetEndAt() != time.Date(1,1,1, 0,0,0 ,0,nil)  {
		q = q.Where("? < ?",model.Columns.SaleRecord.Createdat,filter.GetEndAt())
	}
	if filter.GetCountry() != "" {
		res := model.Vendor{}
		if err := vr.conn.Model(model.Vendor{}).
			Where("?=?",model.Columns.Vendor.Country,filter.GetCountry()).
			Select(&res);err!=nil {
				return nil,err
		}
		q = q.Where("?=?",model.Columns.SaleRecord.Vendorid,res.GetID())
	}
	if filter.GetCity() != "" {
		res := model.Vendor{}
		if err := vr.conn.Model(model.Vendor{}).
			Where("?=?",model.Columns.Vendor.City,filter.GetCity()).
			Select(&res);err!=nil {
			return nil,err
		}
		q = q.Where("?=?",model.Columns.SaleRecord.Vendorid,res.GetID())
	}
	if filter.GetCategory() != 0 {
		res := model.Vendor{}
		if err := vr.conn.Model(model.Vendor{}).
			Where("?=?",model.Columns.Vendor.Category,filter.GetCategory()).
			Select(&res);err!=nil {
			return nil,err
		}
		q = q.Where("?=?",model.Columns.SaleRecord.Vendorid,res.GetID())
	}

	if err := q.Select(&response);err!=nil {
		return nil,err
	}

	return response, nil
}

func (vr *VmtRepository) GetVendor(vendorID int32) domain.IVendor{
	res := model.Vendor{}
	if err := vr.conn.Model(model.Vendor{}).
		Where("?=?",model.Columns.Vendor.ID,vendorID).
		Select(&res);err!=nil{
		return nil
	}

	return &res
}

