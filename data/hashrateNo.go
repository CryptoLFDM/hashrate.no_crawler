package data

var CardsResult *CardsHarvested

type CardsHarvested struct {
	CardHarvested []CardHarvested
}

type CardHarvested struct {
	Name       string `json:"card_name"`
	DualCoin   []DualCoin
	SingleCoin []SingleCoin
}

type DualCoin struct {
	HashratePrimary float64 `json:"hashrate_primary"`
	HashrateAlt     float64 `json:"hashrate_alt"`
	CoinName        string  `json:"coin_name"`
	UnitPrimary     string  `json:"primary_unit"`
	UnitAlt         string  `json:"alt_unit"`
	CoinPrimary     string  `json:"primary_coin"`
	CoinAlt         string  `json:"alt_coin"`
	Conso           int     `json:"conso"`
	Income          float64 `json:"income"`
}

type SingleCoin struct {
	CoinName        string  `json:"coin_name"`
	HashratePrimary float64 `json:"hashrate_primary"`
	UnitPrimary     string  `json:"primary_unit"`
	CoinPrimary     string  `json:"primary_coin"`
	Conso           int     `json:"conso"`
	Income          float64 `json:"income"`
}
