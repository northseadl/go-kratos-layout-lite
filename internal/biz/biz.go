package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go-kratos-layout-lite/internal/data"
)

var ServerProviderSet = wire.NewSet(NewHelloUC)
var TaskProviderSet = wire.NewSet(NewHelloUC)
var WorkerProviderSet = wire.NewSet(NewHelloUC)

type HelloUC struct {
	data *data.Data
	h    *log.Helper
}

func NewHelloUC(data *data.Data, logger log.Logger) *HelloUC {
	return &HelloUC{
		data: data,
		h:    log.NewHelper(logger),
	}
}
