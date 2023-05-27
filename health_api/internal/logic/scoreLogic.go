package logic

import (
	"context"

	"wk-middleground/health_api/internal/svc"
	"wk-middleground/health_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScoreLogic {
	return &ScoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScoreLogic) Score(req *types.ScoreRequest) (resp *types.ScoreResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
