//@Author: wulinlin
//@Description:
//@File:  jwtConfig
//@Version: 1.0.0
//@Date: 2023/03/24 11:11

package config

type JwtObj struct {
	JwtPrefix     string `mapstructure:"jwt_prefix"`
	JwtHeaderKey  string `mapstructure:"jwt_header_key"`
	TokenKey      string `mapstructure:"token_key"`
	JwtExpireHour int64  `mapstructure:"jwt_expire_hour"`
	KsrTokenKey   string `mapstructure:"ksr_token_key"`
}
