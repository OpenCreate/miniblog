// Copyright 2023 Innkeeper Ma Shaodong(马少东). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/opencreate/miniblog.

package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opencreate/miniblog/internal/pkg/errno"
)

type ErrResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		hcode, code, msg := errno.Decode(err)
		c.JSON(hcode, ErrResponse{
			Code:    code,
			Message: msg,
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
