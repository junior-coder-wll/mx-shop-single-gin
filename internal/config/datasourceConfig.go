//@Author: wulinlin
//@Description:
//@File:  base
//@Version: 1.0.0
//@Date: 2023/03/24 11:05

package config

type DatasourceObj struct {
	MysqlCfgInfo string   `mapstructure:"mysql"`
	RedisCfgInfo RedisObj `mapstructure:"redis"`
}

type RedisObj struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Pwd      string `mapstructure:"pwd"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"poolSize"`
}
