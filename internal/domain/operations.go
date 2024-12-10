package domain

const (
	BUY  = "buy"
	SELL = "sell"
)

type Operation struct {
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  int     `json:"quantity"`
}

type Tax struct {
	Tax float64 `json:"tax"`
}
