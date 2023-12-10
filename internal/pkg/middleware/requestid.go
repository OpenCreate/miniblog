// Copyright 2023 Innkeeper Ma Shaodong(马少东). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/opencreate/miniblog.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/opencreate/miniblog/internal/pkg/known"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.Request.Header.Get(known.XRequestIDKey)

		if requestID == "" {
			requestID = uuid.New().String()
		}
		// 将 RequestID 保存在 gin.Context 中，方便后边程序使用
		ctx.Set(known.XRequestIDKey, requestID)
		// 将 RequestID 保存在 HTTP 返回头中，Header 的键为 `X-Request-ID`
		ctx.Writer.Header().Set(known.XRequestIDKey, requestID)
		ctx.Next()
	}
}
