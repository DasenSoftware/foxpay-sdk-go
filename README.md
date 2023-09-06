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

源码下载：[foxpay-sdk-go](https://github.com/KamenSoftware/foxpay-sdk-go)
```bash
go get github.com/KamenSoftware/foxpay-sdk-go
```
## 项目使用


### 使用示例
```go
package main

import (
	"fmt"
	"github.com/KamenSoftware/foxpay-sdk-go/common"
	"github.com/KamenSoftware/foxpay-sdk-go/foxpay"
	"github.com/KamenSoftware/foxpay-sdk-go/util"
)

func main() {
	appid := "YOUR-API-ID"
	url := "YOUR-URL"
	publicKey := "YOUR-PUBLIC-KEY"
	privateKey := "YOUR-PRIVATE-KEY"	
	fp := foxpay.NewFoxPay(
		foxpay.WithFoxPayObjAPPID(appid),
		foxpay.WithFoxPayObjURL(url),
		foxpay.WithFoxPayObjPublicKey(publicKey),
		foxpay.WithFoxPayObjPrivateKey(privateKey),
	)
	//查询资产
	balanceResult, err := fp.GetBalance()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(balanceResult)

	//查询订单
	getOrderResult, err := fp.GetApplicationOrder(common.OrderOrTradeNo{
		//OrderNo: util.StringPtr("004"),
		TradeNo: util.StringPtr("AP2023071310442022925526694"),
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(getOrderResult)

	//关闭订单
	err = fp.CloseApplicationOrder(common.OrderOrTradeNo{
		OrderNo: util.StringPtr("004"),
	})
	if err != nil {
		fmt.Println(err)
	}

	//创建订单
	createOrderResp, err := fp.CreateApplicationOrder(common.OrderRequest{
		Subject:     util.StringPtr("subject001"),
		OrderNo:     util.StringPtr("005"),
		Amount:      util.StringPtr("3.33"),
		NotifyUrl:   util.StringPtr("notifyUrl"),
		RedirectUrl: util.StringPtr("redirectUrl"),
		TimeOut:     util.Int32Ptr(5000),
		Locale:      util.StringPtr(common.Zh_CN),
		Remark:      util.StringPtr("remarkTest"),
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(createOrderResp)

	//提现凭证获取
	transPrePareResult, err := fp.TransPrepare(common.TransPrepareRequest{
		OrderNo:   util.StringPtr("005"),
		Amount:    util.StringPtr("3.33"),
		ToAddress: util.StringPtr("0x3810fe9f57f2f792a1522088c1a62d14cd5b86c4"),
		NotifyUrl: util.StringPtr("notifyUrl"),
		Remark:    util.StringPtr("remarkTest"),
        GasType:   util.StringPtr(common.GAS_TYPE_ACCOUNT_BALANCE),
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(transPrePareResult)

	//提现确认
	transResult, err := fp.Trans(common.TransRequest{
		TransToken: util.StringPtr("2bef669b555e4d12b2cea8368ab0d0dcldatfk"),
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(transResult)

	//提现记录查询
	getTransResult, err := fp.GetTrans(
		common.OrderOrTradeNo{
			OrderNo: util.StringPtr("20230715100001"),
		})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(getTransResult)
}
```