package svc

import (
	"wk-middleground/equipment_srv/equipment"
	"wk-middleground/health_api/internal/config"
	"wk-middleground/score_srv/score"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	ScoreSrv     score.Score
	EquipmentSrv equipment.Equipment
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ScoreSrv:     score.NewScore(zrpc.MustNewClient(c.ScoreSrv)),
		EquipmentSrv: equipment.NewEquipment(zrpc.MustNewClient(c.EquipmentSrv)),
	}
}
