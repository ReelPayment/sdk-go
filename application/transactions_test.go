package application

import (
	"fmt"
	"testing"
)

func TestTransactions_Currency(t *testing.T) {
	res, err := NewTransactions("sffkuubyosc9i63w", "8b4jSWfrolOTYuUwkJbdw9IJUBeVWK3O").Currency(CurrencyRequest{})
	for _, val := range res.Data {
		fmt.Println(val)
	}
	fmt.Println(res, err)
}

func TestTransactions_EntrustPay(t *testing.T) {
	res, _ := NewTransactions("el9q0mzllpjhducy", "VqdqyUvFt5vQcmJm78dbdLMeEszdzT8I").EntrustPay(EntrustRequest{
		OutTradeNo: "jSWfrolOTdsadcYuUwkJbdw9IJUBeV",
		Symbol:     "USD",
		Amount:     "0.01",
		Name:       "Product name",
		Image:      "https://reelpay.com/product.jpg",
	})
	fmt.Println(res)
}

func TestTransactions_Amount(t *testing.T) {
	res, _ := NewTransactions("sffkuubyosc9i63w", "8b4jSWfrolOTYuUwkJbdw9IJUBeVWK3O").Amount(AmountRequest{
		CurrencyID: "7bifjodnuQvsyU7umAiveISacVbBQT7A",
		FiatName:   "USD",
		FiatAmount: "0.000000001",
	})
	fmt.Println(res)
}

func TestTransactions_Pay(t *testing.T) {
	res, err := NewTransactions("sffkuubyosc9i63w", "8b4jSWfrolOTYuUwkJbdw9IJUBeVWK3O").
		Pay(&PayRequest{
			OutTradeNo: "QYXwrw25PzFZPYw83XWC5ManQWH3SZaW",
			CurrencyID: "BeaMChTcRrC0UkKdSvbLLl9he06RuE01",
			FiatName:   "Usd",
			FiatAmount: "1",
		})
	fmt.Println(res, err)
}

func TestTransactions_Transaction(t *testing.T) {
	res, err := NewTransactions("sffkuubyosc9i63w", "8b4jSWfrolOTYuUwkJbdw9IJUBeVWK3O").Transaction(TransactionRequest{
		TradeNo: "fd720940e69e584c9820230425150922",
	})
	fmt.Println(res, err)
}

func TestTransactions_Close(t *testing.T) {
	res, _ := NewTransactions("sffkuubyosc9i63w", "8b4jSWfrolOTYuUwkJbdw9IJUBeVWK3O").Close(TransactionRequest{
		TradeNo: "5ad7aa46bc7b3daf1720230424164751",
	})
	fmt.Println(res)
}

func TestTransactions_Refund(t *testing.T) {
	res, err := NewTransactions("sffkuubyosc9i63w", "8b4jSWfrolOTYuUwkJbdw9IJUBeVWK3O").Refund(RefundRequest{
		TradeNo: "uz54NEW9TVhLSIDHUSYB4aEFupKoE7I8",
		Fuel:    1,
	})
	fmt.Println(res, err)
}

func TestTransactions_OTCEntrustPay(t *testing.T) {
	res, err := NewTransactions("sffkuubyosc9i63w", "8b4jSWfrolOTYuUwkJbdw9IJUBeVWK3O").OTCEntrustPay(OTCEntrustRequest{
		Amount:     "100",
		OutTradeNo: "jSWfrolOTdsadcYuUwkJbdw9IJUBeV",
	})
	fmt.Println(res, err)
}
