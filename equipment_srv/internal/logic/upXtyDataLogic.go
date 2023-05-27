package logic

import (
	"context"
	"fmt"
	"wk-middleground/equipment_srv/equipment"
	"wk-middleground/equipment_srv/internal/model"

	"wk-middleground/equipment_srv/internal/svc"
	__ "wk-middleground/equipment_srv/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpXtyDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpXtyDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpXtyDataLogic {
	return &UpXtyDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 血糖仪数据上报
func (l *UpXtyDataLogic) UpXtyData(in *__.EquipmentRequest) (*__.ResRequest, error) {
	heathData := new(model.HealthData)
	var xty = new(model.XTYData)
	xty.Gluecose = in.Gluecose
	heathData.Data = xty
	fmt.Println(">>>>>>>>>>>>>>", xty)
	l.svcCtx.HealthM.Insert(l.ctx, heathData)
	return &equipment.ResRequest{}, nil
}
