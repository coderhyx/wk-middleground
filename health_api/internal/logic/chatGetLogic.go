package logic

import (
	"context"

	"wk-middleground/health_api/internal/svc"
	"wk-middleground/health_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatGetLogic {
	return &ChatGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatGetLogic) ChatGet(req *types.ChatGetRequest) (resp *types.ChatGetResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
