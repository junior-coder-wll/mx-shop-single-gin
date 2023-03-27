//@Author: wulinlin
//@Description:
//@File:  LoginHandler
//@Version: 1.0.0
//@Date: 2023/03/24 12:01

package naHandler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HelloController(c *gin.Context) {
	zap.L().Error("这里报错啦")
	c.Error(errors.New("老子日你妈"))
}
