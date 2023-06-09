package logic

import (
	"context"

	"wk-middleground/score_srv/internal/svc"
	__ "wk-middleground/score_srv/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *__.UpdateReq) (*__.CreateResp, error) {
	// todo: add your logic here and delete this line

	return &__.CreateResp{}, nil
}
