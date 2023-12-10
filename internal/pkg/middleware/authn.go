// Copyright 2023 Innkeeper Ma Shaodong(马少东). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/opencreate/miniblog.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/opencreate/miniblog/internal/pkg/core"
	"github.com/opencreate/miniblog/internal/pkg/errno"
	"github.com/opencreate/miniblog/internal/pkg/known"
	"github.com/opencreate/miniblog/pkg/token"
)

//认证（Authentication，简称 Authn）：
//授权（Authorization，简称 Authz）

func Authn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, err := token.ParseRequest(ctx)
		if err != nil {
			core.WriteResponse(ctx, errno.ErrTokenInvalid, nil)
			ctx.Abort()

			return
		}

		ctx.Set(known.XUsernameKey, username)
		ctx.Next()
	}
}
