package domain

import "github.com/go-pg/pg"

type repository interface{
	InitRepository()
	GetInstance() *pg.DB
	SaveSalesRecord(records ISalesRecords) error
	GetSalesRecords(filter Filter) (SearchResponse,error)
}
