package data

var CardsJson []Card

type Card struct {
	HashratePrimary float64 `json:"hashrate_primary"`
	Income          float64 `json:"income"`
	HashrateAlt     float64 `json:"hashrate_alt,omitempty"`
	CoinName        string  `json:"coin_name"`
	Timestamp       string  `json:"@timestamp"`
	Card            string  `json:"card"`
	CoinPrimary     string  `json:"primary_coin"`
	UnitPrimary     string  `json:"primary_unit"`
	UnitAlt         string  `json:"alt_unit,omitempty"`
	CoinAlt         string  `json:"alt_coin,omitempty"`
	Conso           int     `json:"conso"`
}
