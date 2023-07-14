package foxpay

import (
	"encoding/json"
	"github.com/DasenSoftware/foxpay-sdk-go/common"
	errors "github.com/DasenSoftware/foxpay-sdk-go/err"
	"github.com/DasenSoftware/foxpay-sdk-go/status"
	"github.com/DasenSoftware/foxpay-sdk-go/util"
	"io"
	"net/http"
)

/*
*@description 关闭订单
 */

func (o *FoxPay) CloseApplicationOrder(no common.OrderOrTradeNo) (*common.Data, error) {
	url := o.URL + "/sdk/application/closeApplicationOrder" //http请求地址
	para := util.StructToMap(no)
	resp, err := util.HttpRequest(url, para, o.PrivateKey, o.APPID, http.MethodPost)
	if err != nil {
		return nil, errors.NewFoxPaySDKError(status.HttpError, err.Error(), "")
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//处理返回值
	var result common.OrderResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, err
	}
	if result.Code == status.SUCCESS {
		return &result.Data, nil
	} else {
		return nil, errors.NewFoxPaySDKError(result.Code, result.Message, "")
	}
}
