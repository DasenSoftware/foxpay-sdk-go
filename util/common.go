package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"sort"
	"strings"
)

/**
* @description 把map的内容按照字典顺序排
* @param m map数据
* @return string 排序后的字符串
**/

func AlphabeticalOrderSort(m map[string]interface{}) string {
	query := make([]string, 0)
	for k, v := range m {
		query = append(query, fmt.Sprintf("%s=%v", k, v))
	}
	sort.Strings(query)
	queryStr := strings.Join(query, "&")
	return queryStr
}

/**
* @description: 按照文档规定的格式 签名再发起post请求
* @param: uri 请求API的地址
* @param: param map 请求的入参
* @param: privateKey 签名的私钥
* @param: appId 商户ID
* @param: method 请求方式 POST/GET...
* @return: http.response 返回的响应体
**/

func HttpRequest(uri string, pa map[string]interface{}, privateKey string, appId string, method string) (*http.Response, error) {
	bytesData, err := json.Marshal(pa)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(bytesData)
	req, err := http.NewRequest(method, uri, reader)
	if err != nil {
		return nil, err
	}
	pri, err := GetPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	//参数按照字典顺序排序
	paramStr := AlphabeticalOrderSort(pa)
	//签名
	signed, err := Sign(pri, paramStr)
	if err != nil {
		return nil, err
	}
	//设置http header
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("app_id", appId)
	req.Header.Add("sign", signed)
	cli := http.Client{}

	//处理http响应
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/**
* @description struct转map
* @param obj
* @return map[string]interface
**/

func StructToMap(obj interface{}) map[string]interface{} {
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()

	data := make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		value := objValue.Field(i)
		types := objType.Field(i)
		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				continue
			}
			data[types.Tag.Get("mapstructure")] = value.Elem().Interface()
		} else {
			data[types.Tag.Get("mapstructure")] = value.Interface()
		}

	}
	return data
}
