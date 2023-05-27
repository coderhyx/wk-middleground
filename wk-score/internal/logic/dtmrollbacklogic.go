package logic

import (
	"context"

	"wk-middleground/wk-score/internal/svc"
	"wk-middleground/wk-score/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DtmRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDtmRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DtmRollbackLogic {
	return &DtmRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DtmRollbackLogic) DtmRollback(in *__.CreateReq) (*__.CreateResp, error) {
	// todo: add your logic here and delete this line

	return &__.CreateResp{}, nil
}
