// Copyright 2024 Innkeeper cheng <wangcheng.public@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ischeng28/miniblog.

package token

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"sync"
	"time"
)

// Config 包括token包的配置选项
type Config struct {
	key         string
	identityKey string
}

// ErrMissingHeader 表示`Authorization`请求头为空
var ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")

var (
	config = Config{"Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5", "identityKey"}
	once   sync.Once
)

// Init 设置包级别的配置config,config会用于本包后面的token签发和解析
func Init(key string, identityKey string) {
	once.Do(func() {
		if key != "" {
			config.key = key
		}
		if identityKey != "" {
			config.identityKey = identityKey
		}
	})
}

// Parse 使用指定的密钥key解析token,解析成功返回token上下文，否则报错
func Parse(tokenString string, key string) (string, error) {
	//	解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//	确保token加密算法是预期的加密算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(key), nil
	})
	//	解析失败
	if err != nil {
		return "", err
	}

	var identityKey string
	//	如果解析成功，从token中取出token的主题
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		identityKey = claims[config.identityKey].(string)
	}
	return identityKey, err
}

// ParseRequest 从请求头中获取令牌，并将其传递给Parse函数以解析令牌
func ParseRequest(c *gin.Context) (string, error) {
	header := c.Request.Header.Get("Authorization")
	if len(header) == 0 {
		return "", ErrMissingHeader
	}
	var t string
	// 从请求头中取出token
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, config.key)
}

// Sign 使用jwtSecret签发token,token的claims中会存放传入的subject
func Sign(identityKey string) (tokenString string, err error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			config.identityKey: identityKey,
			"nbf":              time.Now().Unix(),
			"iat":              time.Now().Unix(),
			"exp":              time.Now().Add(time.Minute * 30).Unix(),
		})

	// 签发token
	tokenString, err = token.SignedString([]byte(config.key))
	return
}
