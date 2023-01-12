package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GoodsTypeModel = (*customGoodsTypeModel)(nil)

type (
	// GoodsTypeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsTypeModel.
	GoodsTypeModel interface {
		goodsTypeModel
	}

	customGoodsTypeModel struct {
		*defaultGoodsTypeModel
	}
)

// NewGoodsTypeModel returns a model for the database table.
func NewGoodsTypeModel(conn sqlx.SqlConn) GoodsTypeModel {
	return &customGoodsTypeModel{
		defaultGoodsTypeModel: newGoodsTypeModel(conn),
	}
}
