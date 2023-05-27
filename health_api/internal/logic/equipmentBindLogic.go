package logic

import (
	"context"
	"fmt"
	"wk-middleground/equipment_srv/equipment"
	"wk-middleground/score_srv/score"

	"github.com/dtm-labs/client/dtmgrpc"

	"wk-middleground/health_api/internal/svc"
	"wk-middleground/health_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EquipmentBindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEquipmentBindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EquipmentBindLogic {
	return &EquipmentBindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var dtmServer = "etcd://127.0.0.1:2379/dtmservice"

func (l *EquipmentBindLogic) EquipmentBind(req *types.EquipmentBindRequest) (resp *types.EquipmentBindResponse, err error) {
	//dtm 事务id
	gid := dtmgrpc.MustGenGid(dtmServer)
	sagaGrpc := dtmgrpc.NewSagaGrpc(dtmServer, gid)

	//积分
	scoreSrv, err := l.svcCtx.Config.ScoreSrv.BuildTarget()
	if err != nil {
		return nil, err
	}
	createScoreReq := &score.CreateReq{UserId: req.UserId, Score: req.Score}
	sagaGrpc.Add(scoreSrv+"/pb.score/dtmCreate", scoreSrv+"/pb.score/dtmRollback", createScoreReq)

	//绑定健康数据
	eqpSrv, err := l.svcCtx.Config.EquipmentSrv.BuildTarget()
	bindScoreReq := &equipment.BindRequest{
		EquipmentId: req.DataId,
		DataId:      "646df90aa6634dfff81d0aa1",
	}
	fmt.Println(">>>>>>>userId", req.UserId)
	sagaGrpc.Add(eqpSrv+"/pb.Equipment/bindEqpData", eqpSrv+"/pb.Equipment/bindEqpData", bindScoreReq)

	//提交事务
	err = sagaGrpc.Submit()
	if err != nil {
		return nil, fmt.Errorf("submit data to  dtm-server err  : %+v \n", err)
	}
	return
}
