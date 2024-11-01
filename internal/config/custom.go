package config

import (
	"github.com/DrLivsey00/url-shortener-svc/internal/config/custom"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Customer interface {
	Custom() custom.Custom
}

type customer struct {
	getter kv.Getter
	once   comfig.Once
}

func NewCustomer(getter kv.Getter) Customer {
	return &customer{
		getter: getter,
	}
}

func (c *customer) Custom() custom.Custom {
	return c.once.Do(func() interface{} {
		var config struct {
			DomainName string `fig:"domain_name,required"`
		}
		if err := figure.Out(&config).From(kv.MustGetStringMap(c.getter, "custom")).Please(); err != nil {
			panic("error getting custom config: " + err.Error())
		}
		custom := custom.New(config.DomainName)
		return custom
	}).(custom.Custom)
}
