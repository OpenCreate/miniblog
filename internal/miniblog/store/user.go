// Copyright 2023 Innkeeper Ma Shaodong(马少东). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/opencreate/miniblog.

package store

import (
	"context"

	"github.com/opencreate/miniblog/internal/pkg/model"
	"gorm.io/gorm"
)

type UserStore interface {
	Create(c context.Context, user *model.UserM) error
}

type users struct {
	db *gorm.DB
}

var _ UserStore = (*users)(nil)

func newUsers(db *gorm.DB) *users {
	return &users{db}
}

func (u users) Create(c context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error
}
