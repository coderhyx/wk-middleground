package logic

import (
	"context"
	"database/sql"
	"fmt"
	"wk-middleground/score_srv/internal/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"wk-middleground/score_srv/internal/svc"
	__ "wk-middleground/score_srv/pb"

	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
)

type DtmCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDtmCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DtmCreateLogic {
	return &DtmCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DtmCreateLogic) DtmCreate(in *__.CreateReq) (*__.CreateResp, error) {
	scoreM, _ := l.svcCtx.ScoreModel.FindByUserId(l.ctx, in.UserId)
	fmt.Println("--------------->", scoreM)
	//创建子事务屏蔽
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil { //重试
		return nil, status.Error(codes.Internal, err.Error())
	}
	//return nil, status.Error(codes.Internal, err.Error())

	if err := barrier.CallWithDB(l.svcCtx.Db, func(tx *sql.Tx) error {
		fmt.Println(">>>>>>>>>userId:", scoreM.Id)
		if scoreM.Id > 0 { //有记录执行更新
			score := new(model.Score)
			score.Id = scoreM.Id
			score.UserId = in.UserId
			score.Score = scoreM.Score + in.Score
			_, err := l.svcCtx.ScoreModel.DtmUpdate(tx, score)
			if err != nil {
				return fmt.Errorf("返回积分失败 err : %v , score:%+v \n", err, score)
			}
		} else { //无记录新增一个积分
			score := new(model.Score)
			score.UserId = in.UserId
			score.Score = in.Score
			_, err = l.svcCtx.ScoreModel.DtmInsert(tx, score)
			if err != nil {
				return fmt.Errorf("积分生成失败 err : %v , score:%+v \n", err, score)
			}
		}
		return nil
		//return errors.New(">>>>>>>>>>>回滚ba。。。。")
	}); err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	//return nil, status.Error(codes.Internal, err.Error())
	return &__.CreateResp{}, nil
}
