package domain

type repository interface{
	SaveSalesRecord(records ISalesRecords) error
}
