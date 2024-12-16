// Copyright 2024 Innkeeper cheng <wangcheng.public@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ischeng28/miniblog.

package miniblog

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ischeng28/miniblog/internal/miniblog/store"
	"github.com/ischeng28/miniblog/internal/pkg/log"
	"github.com/marmotedu/miniblog/pkg/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	//recommendedHomeDir定义放置miniblog服务配置的默认目录
	recommendedHomeDir = "workspace/miniblog/configs"
	//defaultConfigName 指定了miniblog服务的默认配置文件名.
	defaultConfigName = "miniblog"
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
		log.Errorw("Failed to read viper configuration file", "err", err)
	}

	// 打印 viper 当前使用的配置文件，方便 Debug.
	log.Debugw("Using config file", "file", viper.ConfigFileUsed())

}

// logOptions 从 viper 中读取日志配置，构建 `*log.Options` 并返回.
// 注意：`viper.Get<Type>()` 中 key 的名字需要使用 `.` 分割，以跟 YAML 中保持相同的缩进.
func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}

// initStore 读取 db 配置，创建 gorm.DB 实例，并初始化 miniblog store 层.
func initStore() error {
	dbOptions := &db.MySQLOptions{
		Host:                  viper.GetString("db.host"),
		Username:              viper.GetString("db.username"),
		Password:              viper.GetString("db.password"),
		Database:              viper.GetString("db.database"),
		MaxIdleConnections:    viper.GetInt("db.max-idle-connections"),
		MaxOpenConnections:    viper.GetInt("db.max-open-connections"),
		MaxConnectionLifeTime: viper.GetDuration("db.max-connection-life-time"),
		LogLevel:              viper.GetInt("db.log-level"),
	}

	ins, err := db.NewMySQL(dbOptions)
	if err != nil {
		return err
	}

	_ = store.NewStore(ins)

	return nil
}
