//@Author: wulinlin
//@Description:
//@File:  config
//@Version: 1.0.0
//@Date: 2023/03/10 03:32

package config

import (
	"github.com/spf13/viper"

	"mx-shop-single-gin/pkg/tools"
)

type ServerConfig struct {
	Env        string        `mapstructure:"env"`
	Language   string        `mapstructure:"language"`
	Port       int           `mapstructure:"port"`
	Host       string        `mapstructure:"host"`
	Name       string        `mapstructure:"name"`
	PageSize   int           `mapstructure:"page_size"`
	SecretKey  string        `mapstructure:"secret_key"`
	JwtInfo    JwtObj        `mapstructure:"jwt_info"`
	Datasource DatasourceObj `mapstructure:"datasource"`
	LogInfo    LogObj        `mapstructure:"log"`
}

//
// ConfigLoad
//  @Description: 初始化Config
//
func ConfigLoad(configPath string) *ServerConfig {
	if tools.IsBlank(configPath) {
		configPath = "etc/dev.yaml"
	}
	var config ServerConfig
	v := viper.New()
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		panic(any(err))
	}
	if err := v.Unmarshal(&config); err != nil {
		panic(any(err))
	}
	return &config
}
