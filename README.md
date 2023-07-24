<p align="center">
    <a href="https://reelpay.com/" target="_blank" rel="external">
        <img src="https://github.com/ReelPayment/sdk-php/blob/main/logo.png" height="100px">
    </a>
    <h1 align="center">ReelPay SDK-GO</h1>
    <br>
</p>

ReelPay is committed to helping users quickly transition from Web 2.0 to Web 3.0, attract new users, and collect mainstream currency protocol gateways like BTC, ETH, BSC, SOL, and TRX. Moreover, it aims to help safeguard payment security and provide services such as contract receiving, notifications for address receiving, QR Code receiving, and social promotion.

Documentation is at <a href="https://docs.reelpay.com/" target="_blank" rel="external">https://docs.reelpay.com/</a>.


Requirements
------------
go: >=1.16

At least reelpay sdk-go version 1.0.0 is required for all components to work properly.

Installation
------------

```
go get github.com/ReelPayment/sdk-go
```

Configuration
-------------

To use this extension, you have to configure the Connection class in your application configuration:

```go
import (
    reelpay "github.com/ReelPayment/sdk-go"
)
func main()  {
    transaction := reelpay.Transactions{
        AppID: "eqrbntqbi5uqvkpr",
        AppKey: "XhAlbICW10VJnWGruPL0NSnvb6946JDQ",
    }
    res := transaction.EntrustPay(&reelpay.EntrustPay{
        OutTradeNo: "jSWfrolOTdsadcYuUwkJbdw9IJUBeV",
        Symbol:     "USD",
        Amount:     "1.2",
        Name:       "Product name",
        Image:      "https://reelpay.com/product.jpg",
    })
    fmt.Println(res)
}
```

## Where can I get Appid and Appkey?
Register and open a merchant account, go to application <a href="https://merchant.reelpay.com/" target="_blank" rel="external">management</a> after login background, and create and apply for an application to get the corresponding Appied and Appkey.
