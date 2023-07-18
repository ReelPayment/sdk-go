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
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

const (
	AppidHeader     = "X-Appid"
	TimestampHeader = "X-Timestamp"
	SignHeader      = "X-Sign"
)

type Cover struct {
	Timestamp string
	Body      string
	Sign      string
	AppKey    string
}

func NewCover(appKey string, body interface{}) (*Cover, error) {
	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return &Cover{
		Timestamp: fmt.Sprintf("%d", time.Now().Unix()),
		AppKey:    appKey,
		Body:      string(jsonStr),
	}, nil
}

func (c *Cover) ValidateSign(signParam string) error {
	err := c.hmacSHA256Sign()
	if err != nil {
		return fmt.Errorf("Sign error.")
	}
	if c.Sign != signParam {
		return fmt.Errorf("Sign error: correctSign：%s body：%s", c.Sign, c.Body)
	}
	return nil
}

// sign
func (c *Cover) hmacSHA256Sign() error {
	mac := hmac.New(sha256.New, []byte(c.AppKey))
	_, err := mac.Write([]byte(c.Body + c.Timestamp))
	if err != nil {
		return err
	}
	c.Sign = hex.EncodeToString(mac.Sum(nil))
	return nil
}
