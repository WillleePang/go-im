package models

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserTestModel = (*customUserTestModel)(nil)

type (
	// UserTestModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserTestModel.
	UserTestModel interface {
		userTestModel
	}

	customUserTestModel struct {
		*defaultUserTestModel
	}
)

// NewUserTestModel returns a model for the database table.
func NewUserTestModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserTestModel {
	return &customUserTestModel{
		defaultUserTestModel: newUserTestModel(conn, c, opts...),
	}
}
