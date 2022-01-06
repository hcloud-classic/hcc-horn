package config

import "github.com/Terry-Mao/goconf"

var configLocation = "/etc/hcc/horn/horn.conf"

type hornConfig struct {
	RsakeyConfig *goconf.Section
	GrpcConfig   *goconf.Section
	MysqldConfig *goconf.Section
}
