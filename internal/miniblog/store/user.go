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
	Get(c context.Context, username string) (*model.UserM, error)
	Update(c context.Context, user *model.UserM) error
}

type users struct {
	db *gorm.DB
}

var _ UserStore = (*users)(nil)

func newUsers(db *gorm.DB) *users {
	return &users{db}
}

func (u *users) Create(c context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error
}

func (u *users) Get(c context.Context, username string) (*model.UserM, error) {
	var user model.UserM
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *users) Update(c context.Context, user *model.UserM) error {
	return u.db.Save(user).Error
}
