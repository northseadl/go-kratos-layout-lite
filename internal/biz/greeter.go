package biz

import (
	"context"
	"go-kratos-layout-lite/internal/data"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "go-kratos-layout-lite/api/helloworld/v1"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

type GreeterUC struct {
	data *data.Data
	log  *log.Helper
}

func NewGreeterUsecase(data *data.Data, logger log.Logger) *GreeterUC {
	return &GreeterUC{data: data, log: log.NewHelper(logger)}
}

func (uc *GreeterUC) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return &Greeter{
		Hello: "Hi",
	}, nil
}
