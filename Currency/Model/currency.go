package Currency

type Currency struct {
	Country      string  `json:"country"`
	CurrencyName string  `json:"currency_name"`
	Amount       int     `json:"amount"`
	Code         string  `json:"code"`
	Rate         float64 `json:"rate"`
}
