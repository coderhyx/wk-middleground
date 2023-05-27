package logic

import (
	"context"

	"wk-middleground/wk-score/internal/svc"
	"wk-middleground/wk-score/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoreLogic {
	return &GetScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetScoreLogic) GetScore(in *__.GetReq) (*__.GetResp, error) {
	// todo: add your logic here and delete this line

	return &__.GetResp{}, nil
}
