/*
 * This file is part of the ReelPay SDK-GO package.
 *
 * (c) ReelPay <support@reelpay.com>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */
package application

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

type PayRequest struct {
	OutTradeNo string `json:"out_trade_no"`
	CurrencyID string `json:"currency_id"`
	FiatName   string `json:"fiat_name"`
	FiatAmount string `json:"fiat_amount"`
}

type EntrustRequest struct {
	OutTradeNo string `json:"out_trade_no"`
	Symbol     string `json:"symbol"`
	Amount     string `json:"amount"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	Describe   string `json:"describe"`
}

type EntrustResponse struct {
	Response
	Data struct {
		Url        string `json:"url"`
		TimeExpire int64  `json:"time_expire"`
		ExternalID string `json:"trade_no"`
	} `json:"data"`
}

type AmountRequest struct {
	CurrencyID string `json:"currency_id"`
	FiatName   string `json:"fiat_name"`
	FiatAmount string `json:"fiat_amount"`
}

type Refund struct {
	Response
	Data struct {
		MerchantDeduction string `json:"merchant_deduction"`
		ActualReceipt     string `json:"actual_receipt"`
	} `json:"data"`
}

type AmountResponse struct {
	Response
	Data struct {
		Amount       string `json:"amount"`
		ExchangeRate string `json:"exchange_rate"`
	} `json:"data"`
}

type TransactionRequest struct {
	TradeNo string `json:"trade_no"`
}

type RefundRequest struct {
	TradeNo string `json:"trade_no"`
	Fuel    int8   `json:"fuel"`
}

type TransactionResponse struct {
	Response
	Data struct {
		TradeNo            string `json:"trade_no"`
		OutTradeNo         string `json:"out_trade_no,omitempty"`
		Chain              string `json:"chain"`
		Token              string `json:"token"`
		Contract           string `json:"contract"`
		FiatName           string `json:"fiat_name,omitempty"`
		FiatAmount         string `json:"fiat_amount,omitempty"`
		Amount             string `json:"amount"`
		PayAddress         string `json:"pay_address"`
		Status             string `json:"status,omitempty"`
		TimeExpire         int    `json:"time_expire"`
		ChainID            int64  `json:"chain_id"`
		Mode               string `json:"mode"`
		ContractPayAddress string `json:"contract_pay_address,omitempty"`
		Decimal            int8   `json:"decimal"`
	} `json:"data"`
}

type CurrencyRequest struct{}

type CurrencyResponse struct {
	Response
	Data []*Currency `json:"data"`
}

type Currency struct {
	Chain      string `json:"chain"`
	Token      string `json:"token"`
	Contract   string `json:"contract"`
	Protocol   string `json:"protocol"`
	Logo       string `json:"logo"`
	Decimal    int8   `json:"decimal"`
	CurrencyID string `json:"currency_id"`
}
