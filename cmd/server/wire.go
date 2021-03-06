//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go-kratos-layout-lite/internal/biz"
	"go-kratos-layout-lite/internal/conf"
	"go-kratos-layout-lite/internal/data"
	"go-kratos-layout-lite/internal/server"
	"go-kratos-layout-lite/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ServerProviderSet, service.ProviderSet, newApp))
}
