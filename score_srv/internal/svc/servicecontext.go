package svc

import (
	"database/sql"
	"wk-middleground/score_srv/internal/config"
	"wk-middleground/score_srv/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	ScoreModel model.ScoreModel
	Db         *sql.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, _ := sqlx.NewMysql(c.DB.DataSource).RawDB()
	return &ServiceContext{
		Config:     c,
		ScoreModel: model.NewScoreModel(sqlx.NewMysql(c.DB.DataSource)),
		Db:         db,
	}
}
