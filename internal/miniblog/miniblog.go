// Copyright 2023 Innkeeper Ma Shaodong(马少东). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/opencreate/miniblog.

package miniblog

import (
	"encoding/json"
	"fmt"

	"github.com/opencreate/miniblog/internal/pkg/log"
	"github.com/opencreate/miniblog/pkg/version/verflag"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func NewMiniblogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "miniblog",
		Short:        "A good Go practical project",
		Long:         `A good Go practical project, used to create user with basic information.`,
		SilenceUsage: true,
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, args := range args {
				if len(args) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
		// 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			verflag.PrintAndExitIfRequested()

			log.Init(logOptions())
			defer log.Sync()

			return run()
		},
	}

	cobra.OnInitialize(initConfig)
	// Cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.cobrademo.yaml)")
	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用，子命令不可用
	cmd.Flags().BoolP("toggle", "t", false, "help message for toggle")
	// 添加 --version 标志
	verflag.AddFlags(cmd.PersistentFlags())

	return cmd
}

func run() error {
	settting, _ := json.MarshalIndent(viper.AllSettings(), "", "    ")
	log.Infow(string(settting))
	log.Infow(viper.GetString("db.username"))
	return nil
}
