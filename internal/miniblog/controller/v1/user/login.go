// Copyright 2023 Innkeeper Ma Shaodong(马少东). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/opencreate/miniblog.

package user

import (
	"github.com/gin-gonic/gin"
	"github.com/opencreate/miniblog/internal/pkg/core"
	"github.com/opencreate/miniblog/internal/pkg/errno"
	"github.com/opencreate/miniblog/internal/pkg/log"
	v1 "github.com/opencreate/miniblog/pkg/api/v1"
)

func (ctrl *UserController) Login(c *gin.Context) {
	log.C(c).Infow("Login function called")

	var r v1.LoginRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	resp, err := ctrl.b.Users().Login(c, &r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, resp)
}
