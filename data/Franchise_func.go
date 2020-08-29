package model

func (f *Franchise) GetAPIKey() string {
	return f.Apikey
}

func (f *Franchise) GetID() int32 {
	return int32(f.ID)
}

func (f *Franchise) GetName() string {
	return f.Name
}
