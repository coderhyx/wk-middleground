package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	//EquipmentSrv zrpc.RpcClientConf
	ScoreSrv zrpc.RpcClientConf
	//Mqtt         struct {
	//	Broker   string
	//	ClientID string
	//	Password string
	//}
}
