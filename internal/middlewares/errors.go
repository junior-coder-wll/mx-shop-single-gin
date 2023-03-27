//@Author: wulinlin
//@Description: 错误处理中间件
//@File:  errors
//@Version: 1.0.0
//@Date: 2023/03/10 03:38

package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 如果有错误信息
		if len(c.Errors) > 0 {

			// 获取最后一个错误信息
			err := c.Errors.Last().Err
			fmt.Printf("中间件", c.Errors.String(), "\n===\n")
			fmt.Println(err)
			// 封装错误信息
			//var code int
			//var message string
			//switch err.(type) {
			//case *error:
			//	e := err.(*Error)
			//	code = e.Code
			//	message = e.Message
			//case validator.ValidationErrors:
			//	code = http.StatusBadRequest
			//	message = err.(validator.ValidationErrors).Error()
			//default:
			//	code = http.StatusInternalServerError
			//	message = "Internal Server Error"
			//}
			//
			//// 返回JSON格式的错误信息
			//c.JSON(code, Error{Code: code, Message: message})
			//return
		}
	}
}
