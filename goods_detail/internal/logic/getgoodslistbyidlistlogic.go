package logic

import (
	"context"
	"github.com/reation/micro_service_goods_service/config"
	"github.com/reation/micro_service_goods_service/model"
	"strconv"
	"strings"

	"github.com/reation/micro_service_goods_service/goods_detail/internal/svc"
	"github.com/reation/micro_service_goods_service/protoc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsListByIDListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsListByIDListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsListByIDListLogic {
	return &GetGoodsListByIDListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsListByIDListLogic) GetGoodsListByIDList(in *protoc.GetGoodsListByIDListRequest) (*protoc.GetGoodsListByIDListResponse, error) {
	var idList = make([]string, len(in.GetIdList()))
	for k, v := range in.GetIdList() {
		idList[k] = strconv.FormatInt(v.GetId(), 10)
	}

	goodsInfoList, err := l.svcCtx.GoodsInfoModel.GetGoodsListByIds(l.ctx, strings.Join(idList, ","))
	if err == model.ErrNotFound {
		return &protoc.GetGoodsListByIDListResponse{States: config.RETURE_STATES_EMPTY, GoodList: nil}, nil
	}
	if err != nil {
		return &protoc.GetGoodsListByIDListResponse{States: config.RETURE_STATES_ERROR, GoodList: nil}, nil
	}

	var resp = make([]*protoc.GoodsDataInfo, len(*goodsInfoList))
	for k, v := range *goodsInfoList {
		resp[k] = &protoc.GoodsDataInfo{
			Id:           v.Id,
			Title:        v.Title,
			Cover:        v.Cover,
			Picture:      v.Picture,
			Profile:      v.Profile,
			Prices:       v.Prices,
			Detail:       v.Detail,
			BusinessId:   v.BusinessId,
			BusinessName: "嘟嘟嘟超市",
			State:        v.State,
		}
	}

	return &protoc.GetGoodsListByIDListResponse{States: config.RETURE_STATES_NORMAL, GoodList: resp}, nil
}
