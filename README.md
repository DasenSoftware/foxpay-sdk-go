#  FoxPay SDK GO
## 简介

- 此sdk为方便go开发人员对接foxpay平台收银台功能

- 已实现功能

    - 1.创建订单

    - 2.查询订单

    - 3.关闭订单

    - 4.查询资产
  
    - 5.提现凭证获取

    - 6.提现确认

    - 7.提现记录查询


## 版本要求

golang >=1.20


## 安装

### 手动安装

源码下载：[foxpay-sdk-go](https://github.com/DasenSoftware/foxpay-sdk-go)
```bash
go get github.com/DasenSoftware/foxpay-sdk-go
```
## 项目使用


### 使用示例

```go
//实例化foxpay对象
fp := foxpay.NewFoxPay( 
  foxpay.WithFoxPayObjAPPID("7IJNKYVX"), 
  foxpay.WithFoxPayObjURL("http://139.159.184.46:7600"),
  foxpay.WithFoxPayObjPublicKey("MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCn4Q6mbmlSvH6kbeWJERz4rAQeTB/cEShKzgtkrWyaIZqHLgWh5iNdXyME0uaRUutFae1uc+J1yMyVU3cS+K36JUlqThmBHZ6/93KHsRvQ8FAcmBzKB2yVhW4qF0fA71yaWJzgNe93JI/4u3VSfu7tpy3ilPvmZlh6j9z9I+KKkQIDAQAB"),
  foxpay.WithFoxPayObjPrivateKey("MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAK7w5mS8/G7v7xjTAC5S+qx3E0OBFX6kt0kN3l9OHW7qA6cxFEP4DP04qRD3DCR4VvoHcNO0JwFv0Z3hc0afQt6qFGPcWMk9hFMTtB6gQwjOxLSRxDO2u1EVcfp13KzCLagx8/bfPM8+EMHkSXMPncIXmWXoy64r3aJmHOrCAgQnAgMBAAECgYAZjsMyC3Qbpvz41PatTd0mbh0H2ydvQZQvXZHTvZ9KMXEIL4Dk7yvCoND+VAFXoKcgTw76NtMOAC6RELtdIW5Mx+M4p8OZD+DeNqXUrXUpKd1odyagROSsqm03kwqDC+ZKIpU2f2SRST0HCi3ttXXMWqaKz3aeca6IaUPamoDGeQJBAPZHQRXlo1pbf0PJyjTgWkRB2FbZNCps9qoxhFMOOWf98rH/MvZ/Zi6YJMHqJssW+mK0lkMtgeHS5lwTgkGIqo8CQQC12MCRbBSdjL5hqUuFf8hUMQ+UC6yAEziiGAwkD5FhjpY07ylcywdJMH6srH4cFeYHTVfZPKLqk1yT46GheLjpAkAxSq+nL1ATnK6LJc836BOJB9jCATUkrKxuAf0nFni87KHvqFFN7s/H0aHBwhjDmzTAHr7YcTpGtYxvr2Pps+3XAkA7vhFN9X80X5fwh+ka2+dZ2aBvmAI9NZNmlZXvhvnRXkH09BnXtZAYOIl1e1oXKg6fmYZiBWzUukMxBxkD7qB5AkEAlP4x8e2red7uDNKk0iauppUFuZrX/dd3snm2ulAcC/qjacoXhCSF+KTApD9ScgZ7RJ4ZO2zr1lCfW5WbWaVi4Q=="),
)
//查询余额
fp.GetBalance()

//查询订单
//OrderNo ，TradeNo 二选一或者都填
fp.GetApplicationOrder(
	common.OrderOrTradeNo{
      OrderNo: util.StringPtr("004"),
      TradeNo: util.StringPtr("AP2023071310442022925526694"),
  })

//关闭订单
//OrderNo ，TradeNo 二选一或者都填
fp.CloseApplicationOrder(
	common.OrderOrTradeNo{
      OrderNo: util.StringPtr("004"),
      TradeNo: util.StringPtr("AP2023071310442022925526694"),
})

//创建订单
fp.CreateApplicationOrder(
	common.OrderRequest{
      Subject:     "subjecttest",
      OrderNo:     "004",
      Amount:      "3.33",
      NotifyUrl:   util.StringPtr("nnotifyUrl"),
      RedirectUrl: util.StringPtr("redirectUrl"),
      TimeOut:     5000,
      Locale:      common.Zh_CN,
      Remark:      util.StringPtr("remarkTest"),
    })

//提现凭证获取
fp.TransPrepare(common.TransPrepareRequest{
    OrderNo:   "004",
    Amount:    "3.33",
    ToAddress: "THZB25oFbuogkuHq7BrLXyXLcfKFMesNe9",
    NotifyUrl: util.StringPtr("www.notifyurl.com"),
    Remark:    util.StringPtr("remarkkkkk"),
})

//提现确认
fp.Trans(common.TransRequest{
    TransToken: "2bef669b555e4d12b2cea8368ab0d0dcldatfk",
})

//提现记录查询
fp.GetTrans(
  common.OrderOrTradeNo{
      OrderNo: util.StringPtr("003"),
  })
```