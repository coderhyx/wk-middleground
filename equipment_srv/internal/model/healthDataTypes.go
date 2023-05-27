package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HealthData struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Data     any                `bson:"data,omitempty" json:"data,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

// 血糖仪结构体
type XTYData struct {
	Gluecose float64 `json:"gluecose"`
}

// 血压仪结构体
type XYYData struct {
	Spo2 float64 `json:"spo2"`
	Bpm  float64 `json:"bpm"`
	Pi   float64 `json:"pi"`
}
