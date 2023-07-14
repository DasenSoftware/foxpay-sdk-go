package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
)

/**
 *  @Description: sha1withrsa 签名
 *  @param privateKey 私钥
 *  @param content 需要签名的内容
 *  @return sign 签名
 *  @return err 错误
 */

func Sign(privateKey *rsa.PrivateKey, content string) (sign string, err error) {
	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(content))
	hashed := h.Sum(nil)
	var signature []byte
	signature, err = rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, hashed)
	if err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(signature)
	return
}

/**
 *  @Description: 验签
 *  @param publicKey 公钥
 *  @param originalData 原始数据
 *  @param ciphertext 签名
 *  @return ok 验签结果
 *  @return error 错误
 */

func RSAVerify(publicKey *rsa.PublicKey, originalData, ciphertext string) (ok bool, err error) {
	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(originalData))
	digest := h.Sum(nil)
	body, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return false, err
	}
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA1, digest, body)
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
 *  @Description: 私钥字符串转换成对象
 *  @param privateKey 私钥
 *  @return pri 私钥对象
 *  @return error 错误
 */

func GetPrivateKey(privateKey string) (pri *rsa.PrivateKey, err error) {
	block, _ := base64.StdEncoding.DecodeString(privateKey)
	if block == nil {
		return nil, errors.New("pem.Decode err")
	}
	private, err := x509.ParsePKCS8PrivateKey(block)
	if err != nil {
		return nil, err
	}
	return private.(*rsa.PrivateKey), nil
}

/**
 *  @Description: 公钥字符串转换成对象
 *  @param privateKey 公钥
 *  @return pri 公钥对象
 *  @return error 错误
 */

func GetPublicKey(publicKey string) (pub *rsa.PublicKey, err error) {
	block, err := base64.StdEncoding.DecodeString(publicKey)
	if block == nil {
		return nil, errors.New("pem decode err")
	}
	// 将pem格式公钥文件进行反序列化
	p, err := x509.ParsePKIXPublicKey(block)
	if err != nil {
		return nil, errors.New("public key：" + err.Error())
	}
	return p.(*rsa.PublicKey), err
}
