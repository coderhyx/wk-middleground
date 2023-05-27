package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mongo MongoConf
}

type MongoConf struct {
	Uri        string
	Db         string
	Collection string
}
