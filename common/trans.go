package common

// 提现凭证获取 request

type TransPrepareRequest struct {
	OrderNo   *string `json:"order_no,omitempty" mapstructure:"order_no"`
	Amount    *string `json:"amount,omitempty" mapstructure:"amount"`
	ToAddress *string `json:"to_address"  mapstructure:"to_address"` //提现地址
	NotifyUrl *string `json:"notify_url" mapstructure:"notify_url"`
	Remark    *string `json:"remark,omitempty"  mapstructure:"remark"`
}
type TransPrepare struct {
	TransToken *string `json:"trans_token" mapstructure:"trans_token"`
}
type TransPrepareResponse struct {
	Code    int32        `json:"code"`
	Data    TransPrepare `json:"data"`
	Message string       `json:"message"`
}

// 提现确认

type TransRequest struct {
	TransToken *string `json:"trans_token" mapstructure:"trans_token"`
}
type TransResponse struct {
	Code    int32  `json:"code"`
	Data    Trans  `json:"data"`
	Message string `json:"message"`
}
type Trans struct {
	OrderNo    *string `json:"order_no" mapstructure:"order_no"`
	TradeNo    *string `json:"trade_no"  mapstructure:"trade_no"`
	TransToken *string `json:"trans_token"  mapstructure:"trans_token"` // 交易凭证
	Status     *int32  `json:"status" mapstructure:"status"`            //提现状态：1待提现，2处理中， 3提现成功， 4提现失败
	ToAddress  *string `json:"to_address"  mapstructure:"to_address"`   //提现地址
	NotifyUrl  *string `json:"notify_url" mapstructure:"notify_url"`
	Amount     *string `json:"amount"  mapstructure:"amount"`           //数量
	Remark     *string `json:"remark"  mapstructure:"remark"`           //备注
	CreateTime *int64  `json:"create_time"  mapstructure:"create_time"` //创建时间
}

// 提现记录响应

type GetTransResponse struct {
	Code    int32    `json:"code"`
	Data    GetTrans `json:"data"`
	Message string   `json:"message"`
}
type GetTrans struct {
	OrderNo    *string `json:"order_no" mapstructure:"order_no"`
	TradeNo    *string `json:"trade_no"  mapstructure:"trade_no"`
	TransToken *string `json:"trans_token"  mapstructure:"trans_token"` //交易凭证
	Status     *int32  `json:"status" mapstructure:"status"`            //提现状态：1待提现，2处理中， 3提现成功， 4提现失败
	ToAddress  *string `json:"to_address"  mapstructure:"to_address"`   //提现地址
	NotifyUrl  *string `json:"notify_url" mapstructure:"notify_url"`
	Amount     *string `json:"amount"  mapstructure:"amount"`           //数量
	Fee        *string `json:"fee"  mapstructure:"fee"`                 //手续费
	Balance    *string `json:"balance"  mapstructure:"balance"`         //余额
	PayTime    *int64  `json:"pay_time"  mapstructure:"pay_time"`       //支付时间
	TxHash     *string `json:"tx_hash"  mapstructure:"tx_hash"`         //余额
	Remark     *string `json:"remark"  mapstructure:"remark"`           //备注
	CreateTime *int64  `json:"create_time"  mapstructure:"create_time"` //创建时间
}
