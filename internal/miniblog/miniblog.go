// Copyright 2023 Innkeeper Ma Shaodong(马少东). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/opencreate/miniblog.

package miniblog

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opencreate/miniblog/internal/pkg/core"
	"github.com/opencreate/miniblog/internal/pkg/errno"
	"github.com/opencreate/miniblog/internal/pkg/log"
	"github.com/opencreate/miniblog/internal/pkg/middleware"
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
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()
	mws := []gin.HandlerFunc{gin.Recovery(), middleware.NoCache, middleware.Cors, middleware.Secure, middleware.RequestID()}

	g.Use(mws...)

	g.NoRoute(func(ctx *gin.Context) {
		core.WriteResponse(ctx, errno.ErrPageNotFound, nil)
	})

	g.GET("/health", func(ctx *gin.Context) {
		log.C(ctx).Infow("Healthz function called")
		core.WriteResponse(ctx, nil, map[string]string{"status": "ok"})
	})

	httpsrv := http.Server{
		Addr:    viper.GetString("addr"),
		Handler: g,
	}
	log.Infow("Start to listening the incoming requests on http address", "addr", viper.GetString("addr"))
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Infow("Shutting down server ...")

	// 创建 ctx 用于通知服务器 goroutine, 它有 10 秒时间完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 10 秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过 10 秒就超时退出
	if err := httpsrv.Shutdown(ctx); err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
		return err
	}

	log.Infow("Server exiting")
	return nil
}
