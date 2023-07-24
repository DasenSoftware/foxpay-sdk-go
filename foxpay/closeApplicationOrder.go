package foxpay

import (
	"encoding/json"
	"github.com/KamenSoftware/foxpay-sdk-go/common"
	errors "github.com/KamenSoftware/foxpay-sdk-go/err"
	"github.com/KamenSoftware/foxpay-sdk-go/status"
	"github.com/KamenSoftware/foxpay-sdk-go/util"
	"io"
	"net/http"
)

/*
*@description 关闭订单
 */

func (o *FoxPay) CloseApplicationOrder(ot common.OrderOrTradeNo) error {
	url := o.URL + "/sdk/application/closeApplicationOrder" //http请求地址
	para := util.StructToMap(ot)
	resp, err := util.HttpRequest(url, para, o.PrivateKey, o.APPID, http.MethodPost)
	if err != nil {
		return errors.NewFoxPaySDKError(status.HttpError, err.Error(), "")
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//处理返回值
	var result common.CloseOrderResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return err
	}
	if result.Code != status.SUCCESS {
		return errors.NewFoxPaySDKError(result.Code, result.Message, "")
	}
	return nil
}
