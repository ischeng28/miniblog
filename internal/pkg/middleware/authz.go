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
	"github.com/ischeng28/miniblog/internal/pkg/log"
)

// Auther 用来定义授权接口实现
// sub：操作主题，obj：操作对象，act：操作
type Auther interface {
	Authorize(sub, obj, act string) (bool, error)
}

// Authz 是Gin中间件，用来进行请求授权
func Authz(a Auther) gin.HandlerFunc {
	return func(context *gin.Context) {
		sub := context.GetString(known.XUsernameKey)
		obj := context.Request.URL.Path
		act := context.Request.Method

		log.Debugw("Build authorize context", "sub", sub, "obj", obj, "act", act)
		if allowed, _ := a.Authorize(sub, obj, act); !allowed {
			core.WriteResponse(context, errno.ErrUnauthorized, nil)
			context.Abort()
			return
		}
	}
}
