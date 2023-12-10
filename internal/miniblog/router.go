// Copyright 2023 Innkeeper Ma Shaodong(马少东). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/opencreate/miniblog.

package miniblog

import (
	"github.com/gin-gonic/gin"
	"github.com/opencreate/miniblog/internal/miniblog/controller/v1/user"
	"github.com/opencreate/miniblog/internal/miniblog/store"
	"github.com/opencreate/miniblog/internal/pkg/core"
	"github.com/opencreate/miniblog/internal/pkg/errno"
	"github.com/opencreate/miniblog/internal/pkg/log"
	"github.com/opencreate/miniblog/internal/pkg/middleware"
)

func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	uc := user.New(store.S)
	g.POST("/login", uc.Login)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)
			userv1.PUT(":name/change-password", uc.ChangePassword)
			userv1.Use(middleware.Authn())
		}
	}
	return nil
}
