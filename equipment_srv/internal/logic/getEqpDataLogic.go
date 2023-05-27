package logic

import (
	"context"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"

	"wk-middleground/equipment_srv/equipment"
	"wk-middleground/equipment_srv/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
)

type GetEqpDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEqpDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEqpDataLogic {
	return &GetEqpDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 数据查询
func (l *GetEqpDataLogic) GetEqpData(in *equipment.EquipmentRequest) (*equipment.EquipmentResponse, error) {
	typeInfo := reflect.TypeOf(*in)
	valInfo := reflect.ValueOf(*in)
	filter := bson.D{}
	for i := 0; i < typeInfo.NumField(); i++ {
		k := typeInfo.Field(i).Name
		if i >= 3 && k != "UserId" && k != "EquipmentId" {
			v := valInfo.Field(i).Float()
			if v != 0 {
				k = strings.ToLower(k)
				filter = append(filter, bson.E{"data." + k, v})
			}
		}
	}
	res, _ := l.svcCtx.HealthM.FindByAny(l.ctx, filter)
	var eqpData []*equipment.EquipmentData
	for _, v := range res {
		str := fmt.Sprintf("%v", v.Data)
		eqpData = append(eqpData, &equipment.EquipmentData{
			Id:   hex.EncodeToString(v.ID[:]),
			Data: str,
		})
	}
	return &equipment.EquipmentResponse{
		Data: eqpData,
	}, nil
}
