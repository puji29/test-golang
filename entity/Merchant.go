package entity

type Merchant struct {
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}
