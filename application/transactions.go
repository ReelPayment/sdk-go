/*
 * This file is part of the ReelPay SDK-GO package.
 *
 * (c) ReelPay <support@reelpay.com>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */
package application

import (
	"github.com/ReelPayment/sdk-go/httpRequest"
)

type Transactions struct {
	appID  string
	appKey string
}

func NewTransactions(appID, appKey string) *Transactions {
	return &Transactions{appID: appID, appKey: appKey}
}

func (t *Transactions) Pay(req *PayRequest) (TransactionResponse, error) {
	response := TransactionResponse{}
	err := httpRequest.Request("/v1/transactions/pay", t.appID, t.appKey, req, &response)
	return response, err
}

func (t *Transactions) EntrustPay(req EntrustRequest) (EntrustResponse, error) {
	response := EntrustResponse{}
	err := httpRequest.Request("/v1/transactions/entrust", t.appID, t.appKey, req, &response)
	return response, err
}

func (t *Transactions) OTCEntrustPay(req OTCEntrustRequest) (EntrustResponse, error) {
	response := EntrustResponse{}
	err := httpRequest.Request("/v1/transactions/otc/entrust", t.appID, t.appKey, req, &response)
	return response, err
}

func (t *Transactions) Amount(req AmountRequest) (AmountResponse, error) {
	response := AmountResponse{}
	err := httpRequest.Request("/v1/transactions/amount", t.appID, t.appKey, req, &response)
	return response, err
}

func (t *Transactions) Transaction(req TransactionRequest) (TransactionResponse, error) {
	response := TransactionResponse{}
	err := httpRequest.Request("/v1/transactions", t.appID, t.appKey, req, &response)
	return response, err
}

func (t *Transactions) Close(req TransactionRequest) (Response, error) {
	response := Response{}
	err := httpRequest.Request("/v1/transactions/close", t.appID, t.appKey, req, &response)
	return response, err
}
func (t *Transactions) Refund(req RefundRequest) (Response, error) {
	response := Response{}
	err := httpRequest.Request("/v1/transactions/refund", t.appID, t.appKey, req, &response)
	return response, err
}

func (t *Transactions) Currency(req CurrencyRequest) (CurrencyResponse, error) {
	response := CurrencyResponse{}
	err := httpRequest.Request("/v1/transactions/currency", t.appID, t.appKey, req, &response)
	return response, err
}
