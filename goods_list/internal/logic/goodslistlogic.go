package logic

import (
	"context"

	"micro_service_goods_service/goods_list/internal/svc"
	"micro_service_goods_service/protoc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsListLogic {
	return &GoodsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GoodsListLogic) GoodsList(in *protoc.GoodsListRequest) (*protoc.GoodsListResponse, error) {
	// todo: add your logic here and delete this line

	return &protoc.GoodsListResponse{}, nil
}
