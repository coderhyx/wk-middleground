package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ HealthDataModel = (*customHealthDataModel)(nil)

type (
	// HealthDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHealthDataModel.
	HealthDataModel interface {
		healthDataModel
	}

	customHealthDataModel struct {
		*defaultHealthDataModel
	}
)

// NewHealthDataModel returns a model for the mongo.
func NewHealthDataModel(url, db, collection string) HealthDataModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customHealthDataModel{
		defaultHealthDataModel: newDefaultHealthDataModel(conn),
	}
}
