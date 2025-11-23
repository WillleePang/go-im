package logic

import (
	"context"

	"go-im/user/rpc/internal/svc"
	"go-im/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserResp, error) {
	// todo: add your logic here and delete this line
	if u, ok := users[in.Id]; ok {
		return &user.GetUserResp{
			Id:    u.Id,
			Name:  u.Name,
			Mail:  u.Mail,
			Phone: u.Phone,
		}, nil
	}

	return &user.GetUserResp{}, nil
}
