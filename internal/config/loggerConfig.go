//@Author: wulinlin
//@Description:
//@File:  loggerConfigh
//@Version: 1.0.0
//@Date: 2023/03/24 11:28

package config

type LogObj struct {
	Level       string `mapstructure:"level"`
	Path        string `mapstructure:"path"`
	Filename    string `mapstructure:"filename"`
	ErrFilename string `mapstructure:"err_filename"`
	MaxAge      int    `mapstructure:"max_size"`
	MaxSize     int    `mapstructure:"max_age"`
	MaxBackups  int    `mapstructure:"max_backups"`
	Format      string `mapstructure:"format"`
	Compress    bool   `mapstructure:"compress"`
}
