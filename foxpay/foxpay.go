package foxpay

import "github.com/KamenSoftware/foxpay-sdk-go/common"

// 接口定义
type foxPayer interface {
	GetBalance() (*common.BalanceDetail, error)                                   //查询资产
	GetApplicationOrder(ot common.OrderOrTradeNo) (*common.QueryOrderData, error) //查询订单
	CloseApplicationOrder(ot common.OrderOrTradeNo) error                         //关闭订单
	CreateApplicationOrder(args common.OrderRequest) (*common.Data, error)        //创建订单
	TransPrepare(tp common.TransPrepareRequest) (*common.TransPrepare, error)     //提现凭证获取
	Trans(t common.TransRequest) (*common.Trans, error)                           //提现确认
	GetTrans(ot common.OrderOrTradeNo) (*common.GetTrans, error)                  //提现记录查询
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
