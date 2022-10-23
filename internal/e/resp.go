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

func Wrap(status Status, body interface{}) interface{} {
	return Response{
		Code: int(status),
		Msg:  status.String(),
		Body: utils.If(status == Success, body, nil),
	}
}

func Pong(c *gin.Context, status Status, body interface{}) {
	c.Set("err", status)
	c.JSON(http.StatusOK, Wrap(status, body))
}

type Endpoint func(*gin.Context) (Status, interface{})

func PongWrapper(handler Endpoint) func(*gin.Context) {
	return func(c *gin.Context) {
		status, body := handler(c)
		Pong(c, status, body)
	}
}
