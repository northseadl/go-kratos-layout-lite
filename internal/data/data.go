package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go-kratos-layout-lite/internal/conf"
)

var ProviderSet = wire.NewSet(NewData)

type Data struct {
	// TODO wrapped database client
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

type Clean struct {
	functions []func()
}

func (c *Clean) Add(fun ...func()) {
	c.functions = append(c.functions, fun...)
}

func (c *Clean) Build() func() {
	return func() {
		for i := range c.functions {
			c.functions[i]()
		}
	}
}
