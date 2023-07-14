package foxpay

import "github.com/DasenSoftware/foxpay-sdk-go/common"

// 接口定义
type foxPayer interface {
	GetBalance() (*common.BalanceDetail, error)
	GetApplicationOrder(no common.OrderOrTradeNo) (*common.OrderResponse, error)
	CloseApplicationOrder(no common.OrderOrTradeNo) (*common.OrderResponse, error)
	CreateApplicationOrder(args common.OrderRequest) (*common.OrderResponse, error)
}

// FoxPay 对象
type FoxPay struct {
	URL        string
	APPID      string
	PublicKey  string
	PrivateKey string
}

type Options func(f *FoxPay)

func WithFoxPayObjURL(url string) Options {
	return func(f *FoxPay) {
		f.URL = url
	}
}

func WithFoxPayObjAPPID(appId string) Options {
	return func(f *FoxPay) {
		f.APPID = appId
	}
}

func WithFoxPayObjPublicKey(publicKey string) Options {
	return func(f *FoxPay) {
		f.PublicKey = publicKey
	}
}

func WithFoxPayObjPrivateKey(privateKey string) Options {
	return func(f *FoxPay) {
		f.PrivateKey = privateKey
	}
}

// 实例化FoxPay对象

func NewFoxPay(opts ...Options) *FoxPay {
	o := &FoxPay{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}
