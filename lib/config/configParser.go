package config

import (
	"github.com/Terry-Mao/goconf"
	"hcc/horn/lib/logger"
)

var conf = goconf.New()
var config = hornConfig{}
var err error

func parseRsakey() {
	config.RsakeyConfig = conf.Get("rsakey")
	if config.RsakeyConfig == nil {
		logger.Logger.Panicln("no rsakey section")
	}

	Rsakey.PublicKeyFile, err = config.RsakeyConfig.String("public_key_file")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

func parseGrpc() {
	config.GrpcConfig = conf.Get("grpc")
	if config.GrpcConfig == nil {
		logger.Logger.Panicln("no grpc section")
	}

	Grpc.Port, err = config.GrpcConfig.Int("port")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

func parseMysqld() {
	config.MysqldConfig = conf.Get("mysqld")
	if config.MysqldConfig == nil {
		logger.Logger.Panicln("no mysqld section")
	}

	Mysqld.User, err = config.MysqldConfig.String("user")
	if err != nil {
		logger.Logger.Panicln(err)
	}

	Mysqld.Host, err = config.MysqldConfig.String("host")
	if err != nil {
		logger.Logger.Panicln(err)
	}
}

// Init : Parse config file and initialize config structure
func Init() {
	if err = conf.Parse(configLocation); err != nil {
		logger.Logger.Panicln(err)
	}

	parseRsakey()
	parseGrpc()
	parseMysqld()
}
