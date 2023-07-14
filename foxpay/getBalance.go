package foxpay

import (
	"encoding/json"
	"github.com/DasenSoftware/foxpay-sdk-go/common"
	errors "github.com/DasenSoftware/foxpay-sdk-go/err"
	"github.com/DasenSoftware/foxpay-sdk-go/status"
	"github.com/DasenSoftware/foxpay-sdk-go/util"
	"io"
)

/*
@description: 查询订单
@return:common.BalanceResponse 返回的响应体
*/

func (o *FoxPay) GetBalance() (*common.BalanceDetail, error) {
	url := o.URL + "/sdk/application/getBalance" //http请求地址
	param := make(map[string]interface{})
	resp, err := util.HttpRequest(url, param, o.PrivateKey, o.APPID, "GET")
	if err != nil {
		return nil, errors.NewFoxPaySDKError(status.HttpError, err.Error(), "")
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//处理返回值
	var result common.BalanceResponse
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
		//获取公钥对象
		pubKey, err := util.GetPublicKey(o.PublicKey)
		if err != nil {
			return nil, errors.NewFoxPaySDKError(status.PublicKeyError, err.Error(), "")
		}
		//验签
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
