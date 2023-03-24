//@Author: wulinlin
//@Description:
//@File:  logger
//@Version: 1.0.0
//@Date: 2023/03/10 03:38

package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求开始时间
		startTime := time.Now()
		// 获取请求方法和路径
		method := c.Request.Method
		path := c.Request.URL.Path
		// 处理请求
		c.Next()

		// 获取请求结束时间
		endTime := time.Now()

		// 计算请求处理时间
		latency := endTime.Sub(startTime)

		// 获取响应状态码
		statusCode := c.Writer.Status()

		// 构建日志记录器
		// 预设相应公用字段
		logFields := []zapcore.Field{
			zap.String("ip", c.ClientIP()),
			zap.Int("status_code", statusCode),
			zap.String("method", method),
			zap.String("path", path),
			zap.Duration("latency", latency),
		}

		// 判断响应状态码，并添加相应的日志级别
		switch {
		case statusCode >= 500:
			logger.Error("Server Error", logFields...)
		case statusCode >= 400:
			logger.Warn("Client Error", logFields...)
		default:
			logger.Info("Request processed", logFields...)
		}
	}
}
