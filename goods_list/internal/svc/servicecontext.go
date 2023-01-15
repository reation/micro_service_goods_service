package svc

import (
	"fmt"
	"github.com/reation/micro_service_goods_service/goods_list/internal/config"
	"github.com/reation/micro_service_goods_service/model"
	"github.com/reation/micro_service_stock_service/goods_stock_list/getgoodsstocklist"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type GoodsDB struct {
	GoodsInfoModel model.GoodsInfoModel
	GoodsTypeModel model.GoodsTypeModel
}

type ServiceContext struct {
	Config       config.Config
	GoodsModel   GoodsDB
	StockService getgoodsstocklist.GetGoodsStockList
}

func NewServiceContext(c config.Config) *ServiceContext {
	dataSourceUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		c.Mysql.StockTable.User,
		c.Mysql.StockTable.Passwd,
		c.Mysql.StockTable.Addr,
		c.Mysql.StockTable.Port,
		c.Mysql.StockTable.DBName,
	)

	return &ServiceContext{
		Config:       c,
		StockService: getgoodsstocklist.NewGetGoodsStockList(zrpc.MustNewClient(c.StockService)),
		GoodsModel: GoodsDB{
			GoodsInfoModel: model.NewGoodsInfoModel(sqlx.NewMysql(dataSourceUrl)),
			GoodsTypeModel: model.NewGoodsTypeModel(sqlx.NewMysql(dataSourceUrl)),
		},
	}
}
