// Code generated by goctl. DO NOT EDIT!
// Source: goods_list.proto

package goodslist

import (
	"context"

	"micro_service_goods_service/protoc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GoodsListInfo     = protoc.GoodsListInfo
	GoodsListRequest  = protoc.GoodsListRequest
	GoodsListResponse = protoc.GoodsListResponse

	GoodsList interface {
		GoodsList(ctx context.Context, in *GoodsListRequest, opts ...grpc.CallOption) (*GoodsListResponse, error)
	}

	defaultGoodsList struct {
		cli zrpc.Client
	}
)

func NewGoodsList(cli zrpc.Client) GoodsList {
	return &defaultGoodsList{
		cli: cli,
	}
}

func (m *defaultGoodsList) GoodsList(ctx context.Context, in *GoodsListRequest, opts ...grpc.CallOption) (*GoodsListResponse, error) {
	client := protoc.NewGoodsListClient(m.cli.Conn())
	return client.GoodsList(ctx, in, opts...)
}