//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

import (
	"time"
)

type ColumnsFranchise struct {
	Apikey, ID, Name string
}

type ColumnsSaleRecord struct {
	ID, Vendorid, Goodsname, Price, Currency, Createdat string
}

type ColumnsVendor struct {
	ID, Country, City, Line1, Line2, Category, FranchiseID string
}

type ColumnsSt struct {
	Franchise  ColumnsFranchise
	SaleRecord ColumnsSaleRecord
	Vendor     ColumnsVendor
}

var Columns = ColumnsSt{
	Franchise: ColumnsFranchise{
		Apikey: "apikey",
		ID:     "id",
		Name:   "name",
	},
	SaleRecord: ColumnsSaleRecord{
		ID:        "id",
		Vendorid:  "vendorid",
		Goodsname: "goodsname",
		Price:     "price",
		Currency:  "currency",
		Createdat: "createdat",
	},
	Vendor: ColumnsVendor{
		ID:          "id",
		Country:     "country",
		City:        "city",
		Line1:       "line1",
		Line2:       "line2",
		Category:    "category",
		FranchiseID: "franchise_id",
	},
}

type TableFranchise struct {
	Name, Alias string
}

type TableSaleRecord struct {
	Name, Alias string
}

type TableVendor struct {
	Name, Alias string
}

type TablesSt struct {
	Franchise  TableFranchise
	SaleRecord TableSaleRecord
	Vendor     TableVendor
}

var Tables = TablesSt{
	Franchise: TableFranchise{
		Name:  "franchise",
		Alias: "t",
	},
	SaleRecord: TableSaleRecord{
		Name:  "sales_record",
		Alias: "t",
	},
	Vendor: TableVendor{
		Name:  "vendor",
		Alias: "t",
	},
}

type Franchise struct {
	tableName struct{} `pg:"franchise,alias:t"`

	Apikey string `pg:"apikey"`
	ID     int    `pg:"id"`
	Name   string `pg:"name"`
}

type SaleRecord struct {
	tableName struct{} `pg:"sales_record,alias:t"`

	ID        int       `pg:"id"`
	Vendorid  int       `pg:"vendorid"`
	Goodsname string    `pg:"goodsname"`
	Price     float64   `pg:"price"`
	Currency  string    `pg:"currency"`
	Createdat time.Time `pg:"createdat"`
}

type Vendor struct {
	tableName struct{} `pg:"vendor,alias:t"`

	ID          int    `pg:"id"`
	Country     string `pg:"country"`
	City        string `pg:"city"`
	Line1       string `pg:"line1"`
	Line2       string `pg:"line2"`
	Category    int    `pg:"category"`
	FranchiseID int    `pg:"franchise_id"`
}
