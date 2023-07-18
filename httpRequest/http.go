/*
 * This file is part of the ReelPay SDK-GO package.
 *
 * (c) ReelPay <support@reelpay.com>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */
package httpRequest

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

const APIURL = "https://pay.reelpay.com"

var httpCodeMap = map[int]error{
	http.StatusUnauthorized:  errors.New("Unauthorized"),
	http.StatusBadRequest:    errors.New("Bad Request"),
	http.StatusNotAcceptable: errors.New("Not Acceptable"),
}

type HttpClient struct {
	Client  *resty.Client
	Request *resty.Request
}

func New() *HttpClient {
	client := resty.New()
	client.
		SetTimeout(120 * time.Second).
		SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
			return 0, errors.New("quota exceeded")
		})

	return &HttpClient{
		Client:  client,
		Request: client.R(),
	}
}

func Request(path, appID, appKey string, body interface{}, response interface{}) error {
	cover, err := NewCover(appKey, body)
	if err != nil {
		return err
	}
	if err = cover.hmacSHA256Sign(); err != nil {
		return err
	}
	hClient := New()
	hClient.Client.SetHeader("content-type", "application/json")
	hClient.Client.SetHeader(TimestampHeader, cover.Timestamp)
	hClient.Client.SetHeader(AppidHeader, appID)
	hClient.Client.SetHeader(SignHeader, cover.Sign)
	hClient.Request.SetBody(cover.Body)
	res, err := hClient.Request.Post(APIURL + path)
	if err != nil {
		return err
	}
	if err, ok := httpCodeMap[res.StatusCode()]; ok {
		return err
	}
	if err = json.Unmarshal(res.Body(), response); err != nil {
		return err
	}
	return nil
}
