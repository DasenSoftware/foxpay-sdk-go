package status

/*
*
| 状态码 | 说明                 |
| ------ | -------------------- |
| 10000  | 操作成功             |
| 10002  | 系统异常             |
| 61000  | 请求参数对象不能为空 |
| 61002  | 响应签名异常         |
| 61003  | 请求签名异常         |
| 61004  | 参数异常             |
| 61007  | 应用不可用           |
| 61008  | 订单不存在           |
| 61009  | 当前订单不能操作     |
| 61010  | 订单号重复           |
*/

const (
	SUCCESS              = 10000
	HttpError            = 62000 //http请求异常
	PublicKeyError       = 63000 //公钥错误
	SignatureFailedERROR = 63002 //签名异常
)

const (
	SignatureFailed = "Signature verification failed"
)
