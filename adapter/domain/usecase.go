package domain

type UseCase interface{
	InitRepository()
	SaveSalesRecord(record interface{}) error
	ViewRecordData(query interface{}) interface{}
}