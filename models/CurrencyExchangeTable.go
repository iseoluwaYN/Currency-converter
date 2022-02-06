package models

type ExchangeTable struct {
	CurrencyTable map[string]map[string]float64
}

func CreateExchangeTable() *ExchangeTable {
	exchangeTable := new(ExchangeTable)
	exchangeTable.CurrencyTable = make(map[string]map[string]float64)
	exchangeTable.CurrencyTable["NGN"] = map[string]float64{"GHG": 0.017, "KSH": 0.27}
	exchangeTable.CurrencyTable["GHG"] = map[string]float64{"NGN": 60.27, "KSH": 16.46}
	exchangeTable.CurrencyTable["KSH"] = map[string]float64{"NGN": 3.66, "GHS": 0.061}

	return exchangeTable
}