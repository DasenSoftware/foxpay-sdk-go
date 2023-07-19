package common

// 余额相关的响应

type BalanceResponse struct {
	Code    int32         `json:"code"`
	Data    BalanceDetail `json:"data"`
	Message string        `json:"message"`
}

type BalanceDetail struct {
	Erc20Address     *string `json:"erc20_address" mapstructure:"erc20_address"`
	Trc20Address     *string `json:"trc20_address" mapstructure:"trc20_address"`
	Money            *string `json:"money" mapstructure:"money"`
	TransFreezeMoney *string `json:"trans_freeze_money" mapstructure:"trans_freeze_money"' `
}
