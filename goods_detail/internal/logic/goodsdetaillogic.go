package logic

import (
	"context"
	"github.com/reation/micro_service_stock_service/config"
	"github.com/reation/micro_service_stock_service/goods_stock/getgoodsstock"

	"github.com/reation/micro_service_goods_service/goods_detail/internal/svc"
	"github.com/reation/micro_service_goods_service/model"
	"github.com/reation/micro_service_goods_service/protoc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsDetailLogic {
	return &GoodsDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GoodsDetailLogic) GoodsDetail(in *protoc.GoodsDetailRequest) (*protoc.GoodsDetailResponse, error) {
	resp, err := l.svcCtx.GoodsInfoModel.FindOne(l.ctx, in.GetGoodsId())
	if err == model.ErrNotFound {
		return &protoc.GoodsDetailResponse{States: config.GOODS_STOCK_STATE_EMPTY, GoodDetail: nil}, nil
	}
	if err != nil {
		return &protoc.GoodsDetailResponse{States: config.GOODS_STOCK_STATE_ERROR, GoodDetail: nil}, nil
	}
	var goodsNum int64 = 0
	goodsNumService, _ := l.svcCtx.StockService.GetGoodsStockByGoodsID(l.ctx, &getgoodsstock.GetGoodsStockRequest{GoodsID: resp.Id})
	if goodsNumService.GetStates() == config.GOODS_STOCK_STATE_NORMAL {
		goodsNum = goodsNumService.GetGoodsNums()
	}

	data := protoc.GoodsDataInfo{
		Id:           resp.Id,
		Title:        resp.Title,
		Cover:        resp.Cover,
		Picture:      resp.Picture,
		Profile:      resp.Profile,
		Prices:       resp.Prices,
		Detail:       resp.Detail,
		BusinessId:   resp.BusinessId,
		BusinessName: "嘟嘟嘟超市",
		State:        resp.State,
		StockNum:     goodsNum,
	}

	return &protoc.GoodsDetailResponse{States: config.GOODS_STOCK_STATE_NORMAL, GoodDetail: &data}, nil
}
