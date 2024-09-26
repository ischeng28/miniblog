// Copyright 2024 Innkeeper wangc <wangcheng.public@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ischeng28/miniblog.

package miniblog

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

const (
	//recommendedHomeDir定义放置miniblog服务配置的默认目录
	recommendedHomeDir = ".workspace/miniblog"
	//defaultConfigName 指定了miniblog服务的默认配置文件名.
	defaultConfigName = "miniblog.yaml"
)

func initConfig() {
	if cfgFile != "" {
		//	从命令行选项指定的配置文件中读取
		viper.SetConfigFile(cfgFile)
	} else {
		//	查找用户主目录
		homeDir, err := os.UserHomeDir()
		//	如果获取用户主目录失败，打印Error:xxx错误，并退出程序（退出码为1）
		cobra.CheckErr(err)

		// 将用 `$HOME/<recommendedHomeDir>` 目录加入到配置文件的搜索路径中
		viper.AddConfigPath(filepath.Join(homeDir, recommendedHomeDir))

		// 把当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath(".")

		//	设置配置文件格式为yaml
		viper.SetConfigType("yaml")

		//	配置文件名称（没有文件扩展名）
		viper.SetConfigName(defaultConfigName)
	}
	//读取匹配的环境变量
	viper.AutomaticEnv()

	viper.SetEnvPrefix("MINIBLOG")

	// 以下 2 行，将 viper.Get(key) key 字符串中 '.' 和 '-' 替换为 '_'
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// 打印 viper 当前使用的配置文件，方便 Debug.
	fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())

}
