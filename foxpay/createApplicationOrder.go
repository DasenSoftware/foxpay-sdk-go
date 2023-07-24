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

/**
@description: 创建订单
*/

func (o *FoxPay) CreateApplicationOrder(args common.OrderRequest) (*common.Data, error) {
	url := o.URL + "/sdk/application/createApplicationOrder" //http请求地址
	para := util.StructToMap(args)
	//Http请求
	resp, err := util.HttpRequest(url, para, o.PrivateKey, o.APPID, http.MethodPost)
	if err != nil {
		return nil, errors.NewFoxPaySDKError(status.HttpError, err.Error(), "")
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//处理返回值
	var result common.CreateOrderResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, err
	}
	if status.SUCCESS == result.Code {
		//验证返回的响应体header的sign
		respSign := resp.Header.Get("sign")

		//将body里的返回，转成map，用于验签
		m := util.StructToMap(result.Data)
		//按照字典序排列
		parmaStr := util.AlphabeticalOrderSort(m)
		//获取公钥对象 验签
		pubKey, err := util.GetPublicKey(o.PublicKey)
		if err != nil {
			return nil, errors.NewFoxPaySDKError(status.PublicKeyError, err.Error(), "")
		}
		flag, err := util.RSAVerify(pubKey, parmaStr, respSign)
		if err != nil {
			return nil, err
		}
		if !flag {
			return nil, errors.NewFoxPaySDKError(status.SignatureFailedERROR, status.SignatureFailed, "")
		}
	} else {
		return nil, errors.NewFoxPaySDKError(result.Code, result.Message, "")
	}
	return &result.Data, nil
}
