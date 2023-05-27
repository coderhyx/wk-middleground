package logic

import (
	"context"
	"fmt"
	"wk-middleground/equipment_srv/equipment"
	"wk-middleground/equipment_srv/internal/svc"

	"github.com/dtm-labs/client/dtmgrpc"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindEqpDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindEqpDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindEqpDataLogic {
	return &BindEqpDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 绑定设备
func (l *BindEqpDataLogic) BindEqpData(in *equipment.BindRequest) (*equipment.ResRequest, error) {
	//子事务屏蔽
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	//fmt.Println("--------->>>>>>>>.11111")
	if err := barrier.MongoCall(l.svcCtx.MongoDB, func(sc mongo.SessionContext) error {
		upData := bson.M{
			//"EqpId": in.EquipmentId,
			"EqpId": "ccccccccccccc",
		}
		fmt.Println("--------->>>>>>>>.2222", upData)
		l.svcCtx.HealthM.DtmUpdate(sc, sc.Client(), "646df90aa6634dfff81d0aa1", upData)
		return nil
	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	//upData := bson.M{
	//	"EqpId": in.EquipmentId,
	//}
	//l.svcCtx.HealthM.DtmUpdate(l.ctx, l.svcCtx.MongoDB, "646df90aa6634dfff81d0aa1", upData)

	return &equipment.ResRequest{}, nil
}
