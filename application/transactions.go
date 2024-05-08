/*
 * This file is part of the ReelPay SDK-GO package.
 *
 * (c) ReelPay <support@reelpay.com>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */
package application

import "github.com/ReelPayment/sdk-go/httpRequest"

type Transactions struct {
	AppID  string
	AppKey string
}

func NewTransactions(appID, appKey string) *Transactions {
	return &Transactions{AppID: appID, AppKey: appKey}
}

func (t *Transactions) Pay(req *PayRequest) (*TransactionResponse, error) {
	response := new(TransactionResponse)
	err := httpRequest.Request("/v1/transactions/pay", t.AppID, t.AppKey, req, response)
	return response, err
}

func (t *Transactions) EntrustPay(req *EntrustRequest) (*EntrustResponse, error) {
	response := new(EntrustResponse)
	err := httpRequest.Request("/v1/transactions/entrust", t.AppID, t.AppKey, req, response)
	return response, err
}

func (t *Transactions) Amount(req *AmountRequest) (*AmountResponse, error) {
	response := new(AmountResponse)
	err := httpRequest.Request("/v1/transactions/amount", t.AppID, t.AppKey, req, response)
	return response, err
}

func (t *Transactions) Transaction(req *TransactionRequest) (*TransactionResponse, error) {
	response := new(TransactionResponse)
	err := httpRequest.Request("/v1/transactions", t.AppID, t.AppKey, req, response)
	return response, err
}

func (t *Transactions) Close(req *TransactionRequest) (*Response, error) {
	response := new(Response)
	err := httpRequest.Request("/v1/transactions/close", t.AppID, t.AppKey, req, response)
	return response, err
}
func (t *Transactions) Refund(req *RefundRequest) (*Response, error) {
	response := new(Response)
	err := httpRequest.Request("/v1/transactions/refund", t.AppID, t.AppKey, req, response)
	return response, err
}

func (t *Transactions) Currency(req *CurrencyRequest) (*CurrencyResponse, error) {
	response := new(CurrencyResponse)
	err := httpRequest.Request("/v1/transactions/currency", t.AppID, t.AppKey, req, response)
	return response, err
}
