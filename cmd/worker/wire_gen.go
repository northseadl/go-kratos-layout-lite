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
	"go-kratos-layout-lite/internal/worker"
)

// Injectors from wire.go:

func wireWorker(confData *conf.Data, logger log.Logger) (*worker.Worker, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterUC := biz.NewGreeterUsecase(dataData, logger)
	workerWorker := worker.NewWorker(logger, greeterUC)
	return workerWorker, func() {
		cleanup()
	}, nil
}
