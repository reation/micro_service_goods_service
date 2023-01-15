package logic

import (
	"context"
	"github.com/reation/micro_service_goods_service/config"
	"github.com/reation/micro_service_goods_service/goods_list/internal/svc"
	"github.com/reation/micro_service_goods_service/model"
	"github.com/reation/micro_service_goods_service/protoc"
	"github.com/reation/micro_service_stock_service/goods_stock_list/getgoodsstocklist"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type GoodsStockInfo struct {
	goodsID  int64
	goodsNum int64
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
	limit := in.GetLimit()
	if in.GetLimit() < 1 {
		limit = 20
	}

	id := in.GetId()
	if in.GetId() < 1 {
		id = 0
	}

	typeID := in.GetType()
	if in.GetType() < 1 {
		typeID = 0
	}
	states, goodsList, err := l.getGoodsList(typeID, id, limit)
	if err != nil {
		return &protoc.GoodsListResponse{States: states, GoodList: []*protoc.GoodsListInfo{}}, nil
	}
	var resp = make([]*protoc.GoodsListInfo, len(*goodsList))
	var goodsIDList = make([]*getgoodsstocklist.GetGoodsStockIDList, len(*goodsList))
	for k, v := range *goodsList {
		goodsIDList[k] = &getgoodsstocklist.GetGoodsStockIDList{
			GoodsId: v.Id,
		}
	}

	goodsStockList := l.getGoodsStock(goodsIDList)
	if goodsStockList == nil {
		return &protoc.GoodsListResponse{States: states, GoodList: resp}, nil
	}

	for k, v := range *goodsList {
		resp[k] = &protoc.GoodsListInfo{
			Id:           v.Id,
			Title:        v.Title,
			Cover:        v.Cover,
			Prices:       v.Prices,
			BusinessId:   v.BusinessId,
			BusinessName: "暂无",
			StockNum:     goodsStockList[v.Id].goodsNum,
		}
	}

	return &protoc.GoodsListResponse{States: states, GoodList: resp}, nil
}

func (l *GoodsListLogic) getGoodsList(typeID, id, limit int64) (int64, *[]model.GoodsInfo, error) {
	var goodsTypeID = ""
	if typeID != 0 {
		goodsTypeInfo, err := l.svcCtx.GoodsModel.GoodsTypeModel.FindOne(l.ctx, id)
		if err != nil {
			return config.RETURE_STATES_TYPE_EMPTY, &[]model.GoodsInfo{}, err
		}
		if goodsTypeInfo.ParentId == 0 {
			goodsTypeList, err := l.svcCtx.GoodsModel.GoodsTypeModel.GetTypeSonIDListByParentID(l.ctx, id)
			if err != nil && len(*goodsTypeList) < 1 {
				return config.RETURE_STATES_SECOND_EMPTY, &[]model.GoodsInfo{}, err
			}

			var typeIDList = make([]string, len(*goodsTypeList))
			for k, v := range *goodsTypeList {
				typeIDList[k] = strconv.FormatInt(v.Id, 10)
			}
			goodsTypeID = strings.Join(typeIDList, ",")

		} else {
			goodsTypeID = strconv.FormatInt(goodsTypeInfo.Id, 10)
		}
	}
	goodsListData, err := l.svcCtx.GoodsModel.GoodsInfoModel.GetGoodsListInTypeID(l.ctx, goodsTypeID, id, limit)
	if err != nil {
		return config.RETURE_STATES_ERROR, &[]model.GoodsInfo{}, err
	}
	if err == model.ErrNotFound {
		return config.RETURE_STATES_EMPTY, &[]model.GoodsInfo{}, err
	}

	return config.RETURE_STATES_NORMAL, goodsListData, nil
}

func (l *GoodsListLogic) getGoodsStock(idList []*getgoodsstocklist.GetGoodsStockIDList) map[int64]GoodsStockInfo {
	goodsStockListResponse, err := l.svcCtx.StockService.GetGoodsStockByGoodsIDList(l.ctx, &getgoodsstocklist.GetGoodsStockListRequest{GoodsIDList: idList})

	if goodsStockListResponse.GetStates() != config.RETURE_STATES_NORMAL || err != nil {
		return nil
	}
	var resp = make(map[int64]GoodsStockInfo)
	for _, v := range goodsStockListResponse.GetGoodsStockList() {
		resp[v.GetGoodsId()] = GoodsStockInfo{
			goodsID:  v.GetGoodsId(),
			goodsNum: v.GetGoodsNum(),
		}
	}

	return resp
}
