package logic

import (
	"context"

	"wk-middleground/health_api/internal/svc"
	"wk-middleground/health_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatisticsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatisticsLogic {
	return &StatisticsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatisticsLogic) Statistics(req *types.StatisticsRequest) (resp *types.StatisticsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
