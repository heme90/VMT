package _type

type VendorCategory int16

const (
	UNKNOWN VendorCategory = 0
	DRINK VendorCategory = iota + 1
	FOOD
)


func (v VendorCategory) Int() int16 {
	return int16(v)
}

var VendorCategoryStringMap = map[string]int16{
	"UNKNOWN": 0,
	"DRINK" : 1,
	"FOOD" : 2,
}

var VendorCategoryCodeMap = map[int16]string{
	0 : "UNKNOWN",
	1 : "DRINK",
	2 : "FOOD",
}