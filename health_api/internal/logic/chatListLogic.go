package logic

import (
	"context"

	"wk-middleground/health_api/internal/svc"
	"wk-middleground/health_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatListLogic {
	return &ChatListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatListLogic) ChatList(req *types.ChatListRequest) (resp *types.ChatListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
