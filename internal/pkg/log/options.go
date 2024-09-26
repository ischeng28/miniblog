// Copyright 2024 Innkeeper wangc <wangcheng.public@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ischeng28/miniblog.

package log

import "go.uber.org/zap/zapcore"

// Options包含与日志中显示调用日志所在的文件与行号
type Options struct {
	//	是否开启caller，如果开启会在日志中显示调用日志所在的文件与行号
	DisableCaller bool
	//	是否禁止在panic及以上级别打印堆栈信息
	DisableStacktrace bool
	//指定日志级别，可选值：debug,info,warn,error,dpanic,panic,fatal
	Level string
	//指定日志显示格式,可选值:console,json
	Format string
	//指定日志输出位置
	OutputPaths []string
}

// NewOptions 创建一个带有默认参数的Options对象
func NewOptions() *Options {
	return &Options{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             zapcore.InfoLevel.String(),
		Format:            "console",
		OutputPaths:       []string{"stdout"},
	}
}
