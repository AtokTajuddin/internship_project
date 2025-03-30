package model

// Province represents an Indonesian province
type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Regency represents an Indonesian regency/city
type Regency struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

// District represents an Indonesian district
type District struct {
	ID        string `json:"id"`
	RegencyID string `json:"regency_id"`
	Name      string `json:"name"`
}

// Village represents an Indonesian village
type Village struct {
	ID         string `json:"id"`
	DistrictID string `json:"district_id"`
	Name       string `json:"name"`
}
