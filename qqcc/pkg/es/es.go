package es

// @Description
// @Author 代码小学生王木木

import (
	"github.com/olivere/elastic/v7"
	"time"
)

type (
	Config struct {
		Url   string `mapstructure:"urls"`
		Sniff bool   `mapstructure:"sniff"`
	}
)

func NewEs(cfg *Config) (*elastic.Client, error) {
	const timeout = 100 * time.Second
	opts := []elastic.ClientOptionFunc{
		elastic.SetURL(cfg.Url),
		elastic.SetSniff(cfg.Sniff),
		elastic.SetHealthcheckTimeoutStartup(timeout),
	}
	return elastic.NewClient(opts...)
}

func MustNewEs(cfg *Config) *elastic.Client {
	es, err := NewEs(cfg)
	if err != nil {
		panic(err)
	}
	return es
}
