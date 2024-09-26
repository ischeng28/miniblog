// Copyright 2024 Innkeeper wangc <wangcheng.public@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ischeng28/miniblog.

package miniblog

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		//	指定命令的名字，该名字会出现在帮助信息里
		Use: "miniblog",
		//命令的简短描述
		Short: "A good Go practical project",
		//命令的详细描述
		Long: `A good Go practical project, used to create user with basic information.

Find more miniblog information at:
	https://github.com/marmotedu/miniblog#readme`,
		//命令出错时，不打印帮助信息。不需要打印帮助信息，设置为true可以保持命令出错时一眼就看到错误信息
		SilenceUsage: true,
		//这里设置命令运行时，不需要指定命令行参数
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		//指令调用cmd.Execute()时，执行的Run函数，函数执行失败会返回错误信息
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments,got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	//以下设置，使得initConfig函数在每个命令运行时都会被调用以读取配置
	cobra.OnInitialize(initConfig)
	//在这里定义标志和配置设置

	//Cobra支持持久化标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的子命令
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the miniblog configuration file.Empty string for no configuration file.")

	//Cobra也支持本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return cmd
}

// run 函数是实际的业务代码入口函数
func run() error {
	//打印所有的配置项及其值
	settings, _ := json.Marshal(viper.AllSettings())
	fmt.Println(string(settings))
	//打印 db->username配置项的值
	fmt.Println(viper.GetString("db.username"))
	return nil
}
