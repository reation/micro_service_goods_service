package svc

import (
	"fmt"
	"github.com/reation/micro_service_goods_service/goods_detail/internal/config"
	"github.com/reation/micro_service_goods_service/model"
	"github.com/reation/micro_service_stock_service/goods_stock/getgoodsstock"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	GoodsInfoModel model.GoodsInfoModel
	StockService   getgoodsstock.GetGoodsStock
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
		Config:         c,
		StockService:   getgoodsstock.NewGetGoodsStock(zrpc.MustNewClient(c.StockService)),
		GoodsInfoModel: model.NewGoodsInfoModel(sqlx.NewMysql(dataSourceUrl)),
	}
}
