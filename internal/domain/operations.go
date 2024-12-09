package domain

const (
	BUY  = "buy"
	SELL = "sell"
)

type Operation struct {
	Operation string  `json:"operation"`
	UnitCost  float32 `json:"unit-cost"`
	Quantity  int     `json:"quantity"`
}

type Tax struct {
	Tax float32 `json:"tax"`
}
