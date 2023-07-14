package common

// 查询订单或者关闭订单的请求参数
type OrderOrTradeNo struct {
	OrderNo *string `json:"order_no,omitempty" mapstructure:"order_no"`
	TradeNo *string `json:"trade_no,omitempty" mapstructure:"trade_no"`
}

// 创建订单的结构体
type OrderRequest struct {
	Subject     string  `json:"subject"  mapstructure:"subject"`
	OrderNo     string  `json:"order_no" mapstructure:"order_no"`
	Amount      string  `json:"amount"   mapstructure:"amount"`
	NotifyUrl   *string `json:"notify_url,omitempty"   mapstructure:"notify_url"`
	RedirectUrl *string `json:"redirect_url,omitempty" mapstructure:"redirect_url"`
	TimeOut     int32   `json:"time_out"  mapstructure:"time_out"`
	Locale      string  `json:"locale"  mapstructure:"locale"`
	Remark      *string `json:"remark,omitempty"  mapstructure:"remark"`
}

// 订单相关的响应
type OrderResponse struct {
	Code    int32  `json:"code" mapstructure:"code"`
	Data    Data   `json:"data" mapstructure:"data"`
	Message string `json:"message"mapstructure:"message"`
}

type Data struct {
	OrderNo     string `json:"order_no" mapstructure:"order_no"`
	Amount      string `json:"amount"  mapstructure:"amount"`
	CreateTime  int64  `json:"create_time"  mapstructure:"create_time"`
	Subject     string `json:"subject"  mapstructure:"subject"`
	ExpireTime  int64  `json:"expire_time"  mapstructure:"expire_time"`
	Remark      string `json:"remark"  mapstructure:"remark"`
	Locale      string `json:"locale"  mapstructure:"locale"`
	NotifyUrl   string `json:"notify_url"  mapstructure:"notify_url"`
	TimeOut     int32  `json:"time_out"  mapstructure:"time_out"`
	TronAddress string `json:"tron_address"  mapstructure:"tron_address"`
	EthAddress  string `json:"eth_address"  mapstructure:"eth_address"`
	TradeNo     string `json:"trade_no"  mapstructure:"trade_no"`
	Status      int32  `json:"status"  mapstructure:"status"`
	PayUrl      string `json:"pay_url"  mapstructure:"pay_url"`
	RedirectUrl string `json:"redirect_url"  mapstructure:"redirect_url"`
}
