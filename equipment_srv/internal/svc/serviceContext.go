package svc

import (
	"context"
	"wk-middleground/equipment_srv/internal/config"
	"wk-middleground/equipment_srv/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceContext struct {
	Config  config.Config
	HealthM model.HealthDataModel
	MongoDB *mongo.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	clientOptions := options.Client().ApplyURI(c.Mongo.Uri)
	MongoDBClient, _ := mongo.Connect(context.TODO(), clientOptions)

	return &ServiceContext{
		Config:  c,
		HealthM: model.NewHealthDataModel(c.Mongo.Uri, c.Mongo.Db, c.Mongo.Collection),
		MongoDB: MongoDBClient,
	}
}
