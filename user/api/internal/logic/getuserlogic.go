// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"go-im/user/api/internal/svc"
	"go-im/user/api/internal/types"
	"go-im/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (resp *types.GetUserResp, err error) {
	// todo: add your logic here and delete this line
	getUserResp, err := l.svcCtx.User.GetUser(l.ctx, &user.GetUserReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetUserResp{
		Id:    getUserResp.Id,
		Mail:  getUserResp.Mail,
		Name:  getUserResp.Name,
		Phone: getUserResp.Phone,
	}, nil
}
