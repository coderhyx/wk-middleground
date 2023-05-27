package logic

import (
	"context"
	"fmt"
	"wk-middleground/equipment_srv/equipment"
	"wk-middleground/equipment_srv/internal/model"
	__ "wk-middleground/equipment_srv/pb"

	"wk-middleground/equipment_srv/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpXyyDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpXyyDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpXyyDataLogic {
	return &UpXyyDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 血氧仪数据上报
func (l *UpXyyDataLogic) UpXyyData(in *__.EquipmentRequest) (*__.ResRequest, error) {
	heathData := new(model.HealthData)
	var xyy = new(model.XYYData)
	xyy.Spo2 = in.Spo2
	xyy.Bpm = in.Bpm
	xyy.Pi = in.Pi
	heathData.Data = xyy
	fmt.Println(">>>>>>>>>>>>>>", heathData.Data)
	l.svcCtx.HealthM.Insert(l.ctx, heathData)
	return &equipment.ResRequest{}, nil
}
