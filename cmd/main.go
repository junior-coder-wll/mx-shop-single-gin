//@Author: wulinlin
//@Description: 项目启动文件
//@File:  main
//@Version: 1.0.0
//@Date: 2023/03/24 10:37

package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"mx-shop-single-gin/internal/config"
	"mx-shop-single-gin/internal/routers"
	"mx-shop-single-gin/pkg/app/mxLog"
)

var configFilePath = flag.String("config", "etc/dev.yaml", "配置文件路径")

func main() {
	// 1. 初始化配置文件
	serverCfg := config.ConfigLoad(*configFilePath)
	fmt.Println(serverCfg)
	// 2. 初始化日志配置
	loggerCfg := mxLog.LogConfig{
		Filename:    serverCfg.LogInfo.Filename,
		ErrFilename: serverCfg.LogInfo.ErrFilename,
		MaxAge:      serverCfg.LogInfo.MaxAge,
		MaxSize:     serverCfg.LogInfo.MaxSize,
		MaxBackups:  serverCfg.LogInfo.MaxBackups,
		Format:      serverCfg.LogInfo.Format,
		Compress:    serverCfg.LogInfo.Compress,
	}
	logger, err := mxLog.NewLogger(true, &loggerCfg)
	if err != nil {

	}
	// 3. 初始化资源配置

	engine := routers.InitRouter(logger)
	zap.L().Error("傻逼狗日的")
	engine.Run(":8023")

}

//func init() {
//	// 1. 初始化配置文件
//	serverCfg := config.ConfigLoad(*configFilePath)
//	fmt.Println(serverCfg)
//	// 2. 初始化日志配置
//
//	// 3. 初始化资源配置
//
//}
