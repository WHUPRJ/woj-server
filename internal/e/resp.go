package e

import (
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Body interface{} `json:"body"`
}

func Wrap(err Err, body interface{}) interface{} {
	return Response{
		Code: int(err),
		Msg:  err.String(),
		Body: utils.If(err == Success, body, nil),
	}
}

func Pong(c *gin.Context, err Err, body interface{}) {
	c.JSON(http.StatusOK, Wrap(err, body))
}
