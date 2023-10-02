package entity

type Cep struct {
	Cep       string `json:"code"`
	State     string `json:"state"`
	City      string `json:"city"`
	Districit string `json:"districit"`
	Address   string `json:"address"`
}
