// Copyright 2024 Innkeeper cheng <wangcheng.public@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ischeng28/miniblog.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ischeng28/miniblog/internal/pkg/core"
	"github.com/ischeng28/miniblog/internal/pkg/errno"
	"github.com/ischeng28/miniblog/internal/pkg/known"
	"github.com/ischeng28/miniblog/pkg/token"
)

func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		//	 解析jwt token
		username, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Set(known.XUsernameKey, username)
		c.Next()
	}
}
