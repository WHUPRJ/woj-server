package problem

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/gin-gonic/gin"
)

type detailsRequest struct {
	Pid uint `form:"pid"`
}

// Details
// @Summary     get details of a problem
// @Description get details of a problem
// @Accept      application/x-www-form-urlencoded
// @Produce     json
// @Param       pid formData int true "problem id"
// @Response    200 {object} e.Response "problem details"
// @Router      /v1/problem/details [post]
func (h *handler) Details(c *gin.Context) {
	req := new(detailsRequest)

	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, err.Error())
		return
	}

	claim, exist := c.Get("claim")
	shouldEnable := !exist || claim.(*global.Claim).Role < model.RoleAdmin

	p, status := h.problemService.Query(req.Pid, true, shouldEnable)
	if status != e.Success {
		e.Pong(c, status, nil)
		return
	}

	pv, status := h.problemService.QueryLatestVersion(req.Pid)
	e.Pong(c, status, gin.H{
		"problem": p,
		"context": pv.Context,
	})
	return
}
