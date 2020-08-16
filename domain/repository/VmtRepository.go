package repository

import (
	"VendingMachineTracker/adapter/domain"
	"errors"
	"github.com/go-pg/pg"
	"sync"
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
	if _,err := tx.Model().Insert(record); err!= nil {
		err := tx.Rollback()
		if err != nil {
			return errors.New("error in rollback transaction!")
		}
		return errors.New("error in save record : ")
	}

	return nil
}

