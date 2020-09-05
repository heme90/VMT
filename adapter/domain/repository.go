package domain

import "github.com/go-pg/pg"

type Repository interface{
	InitRepository() Repository
	GetInstance() *pg.DB
	SaveSalesRecord(records ISalesRecords) error
	SearchSalesRecords(filter Filter) (SearchResponse,error)
	GetFranchiseAuth(apiKey string) int32
}

