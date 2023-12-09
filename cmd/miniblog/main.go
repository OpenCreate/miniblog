// Copyright 2023 Innkeeper Ma Shaodong(马少东). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/opencreate/miniblog.

package main

import (
	"os"

	"github.com/opencreate/miniblog/internal/miniblog"
	_ "go.uber.org/automaxprocs/maxprocs"
)

func main() {
	command := miniblog.NewMiniblogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
