package app

import (
	"context"

	"PowerX/internal/svc"
	"PowerX/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWeWorkAppGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListWeWorkAppGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWeWorkAppGroupLogic {
	return &ListWeWorkAppGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//
// ListWeWorkAppGroup
//  @Description:
//  @receiver this
//  @param req
//  @return resp
//  @return err
//
func (this *ListWeWorkAppGroupLogic) ListWeWorkAppGroup(req *types.AppGroupListRequest) (resp *types.AppGroupListReply, err error) {

	var chatIds []string
	if req.ChatID != `` {
		chatIds = append(chatIds, req.ChatID)
	}
	replys, err := this.svcCtx.PowerX.SCRM.Wechat.PullListWeWorkAppGroupRequest(chatIds...)

	return &types.AppGroupListReply{
		replys,
	}, err
}
