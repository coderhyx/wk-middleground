package logic

import (
	"context"

	"wk-middleground/score_srv/internal/svc"
	__ "wk-middleground/score_srv/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *__.CreateReq) (*__.CreateResp, error) {
	// todo: add your logic here and delete this line

	return &__.CreateResp{}, nil
}
