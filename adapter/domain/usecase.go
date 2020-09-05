package domain

type UseCase interface{
	InitRepository()
	SaveSalesRecord(record ISalesRecords) error
	ViewRecordData(searchFilter Filter) SearchResponse
	AuthFranchiseAdmin(apikey string) (int32,error)
}