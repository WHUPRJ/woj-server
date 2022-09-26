package problem

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/gin-gonic/gin"
)

type searchRequest struct {
	Pid    uint   `form:"pid"`
	Search string `form:"search"`
}

// Search
// @Summary     get detail of a problem
// @Description get detail of a problem
// @Accept      application/x-www-form-urlencoded
// @Produce     json
// @Param       pid formData int false "problem id"
// @Param       search formData string false "search problem"
// @Response    200 {object} e.Response "problem info"
// @Router      /v1/problem/search [post]
func (h *handler) Search(c *gin.Context) {
	req := new(searchRequest)

	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, err.Error())
		return
	}

	if req.Pid == 0 && req.Search == "" {
		e.Pong(c, e.InvalidParameter, nil)
		return
	}

	if req.Pid != 0 {
		problem, status := h.problemService.Query(req.Pid)
		e.Pong(c, status, problem)
		return
	} else {
		problem, status := h.problemService.QueryFuzz(req.Search)
		e.Pong(c, status, problem)
		return
	}
}
