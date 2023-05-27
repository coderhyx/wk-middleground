package svc

import (
	"wk-middleground/health_api/internal/config"
	"wk-middleground/score_srv/score"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	ScoreSrv score.Score
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		ScoreSrv: score.NewScore(zrpc.MustNewClient(c.ScoreSrv)),
	}
}
