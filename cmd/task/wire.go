//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go-kratos-layout-lite/internal/biz"
	"go-kratos-layout-lite/internal/conf"
	"go-kratos-layout-lite/internal/data"
	"go-kratos-layout-lite/internal/task"
)

func wireTasker(*conf.Data, log.Logger) (*task.Tasker, func(), error) {
	panic(wire.Build(data.ProviderSet, biz.TaskProviderSet, task.ProviderSet))
}
