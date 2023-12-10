// Copyright 2023 Innkeeper Ma Shaodong(马少东). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/opencreate/miniblog.

package user

import (
	"github.com/opencreate/miniblog/internal/miniblog/biz"
	"github.com/opencreate/miniblog/internal/miniblog/store"
)

type UserController struct {
	b biz.IBiz
}

func New(ds store.IStore) *UserController {
	return &UserController{b: biz.NewBiz(ds)}
}
