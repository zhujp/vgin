package response

import (
	"github.com/gin-gonic/gin"
)

// type Gin struct {
// 	Ctx *gin.Context
// }

// type Response struct {
// 	Code    int         `json:"code"`
// 	Message string      `json:"msg"`
// 	Data    interface{} `json:"data"`
// }

// func (g *Gin) Response(code int, msg string, data interface{}) {
// 	g.Ctx.JSON(200, Response{
// 		Code:    code,
// 		Message: msg,
// 		Data:    data,
// 	})
// 	return
// }

func Response(c *gin.Context, code int, msg string, json interface{}, err string) {

	if json == nil {
		json = gin.H{}
	}
	c.JSON(200, gin.H{
		"code":  code,
		"msg":   msg,
		"data":  json,
		"error": err,
	})
}
