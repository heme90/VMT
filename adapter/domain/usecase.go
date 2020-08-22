package domain

type UseCase interface{
	InitRepository()
	SaveSalesRecord(record ISalesRecords) error
	ViewRecordData(searchFilter Filter) SearchResponse
	AuthFranchiseAdmin(apikey string) error
}