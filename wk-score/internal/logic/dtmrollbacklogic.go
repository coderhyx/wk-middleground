package logic

import (
	"context"
	"fmt"

	"wk-middleground/wk-score/internal/svc"
	__ "wk-middleground/wk-score/pb"

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
	fmt.Println(">>>>>>>>>>>>>>>>>>>>score  DtmRollback")
	return &__.CreateResp{}, nil
}
