package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GoodsInfoModel = (*customGoodsInfoModel)(nil)

type (
	// GoodsInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsInfoModel.
	GoodsInfoModel interface {
		goodsInfoModel
	}

	customGoodsInfoModel struct {
		*defaultGoodsInfoModel
	}
)

// NewGoodsInfoModel returns a model for the database table.
func NewGoodsInfoModel(conn sqlx.SqlConn) GoodsInfoModel {
	return &customGoodsInfoModel{
		defaultGoodsInfoModel: newGoodsInfoModel(conn),
	}
}
