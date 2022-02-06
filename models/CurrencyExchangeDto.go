package models

type CurrencyExchangeRequestDto struct {
	AmountToConvert float64
	From string
	To string
}

type CurrencyExchangeResponseDto struct {
	AmountToConvert float64
	From string
	To string
	AmountConverted float64
}
