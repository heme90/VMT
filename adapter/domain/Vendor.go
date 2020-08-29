package domain

type IVendor interface{
	GetID() int32
	GetLocation() IAddress
	GetCategory() int16
	GetFranchiseID() int32
}


type IAddress interface {
	GetCountry() string
	GetCity() string
	GetLine1() string
	GetLine2() string
}

type IFranchise interface{
	GetAPIKey() string
	GetID() int32
	GetName() string
}