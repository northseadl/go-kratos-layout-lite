// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"go-kratos-layout-lite/internal/biz"
	"go-kratos-layout-lite/internal/conf"
	"go-kratos-layout-lite/internal/data"
	"go-kratos-layout-lite/internal/task"
)

// Injectors from wire.go:

func wireTasker(confData *conf.Data, logger log.Logger) (*task.Tasker, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	helloUC := biz.NewHelloUC(dataData, logger)
	tasker := task.NewTasker(logger, helloUC)
	return tasker, func() {
		cleanup()
	}, nil
}
