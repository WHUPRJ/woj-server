package status

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/gin-gonic/gin"
)

type queryRequest struct {
	SubmissionID uint `form:"sid"`
}

func (h *handler) Query(c *gin.Context) {
	req := new(queryRequest)

	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, err.Error())
		return
	}

	status, eStatus := h.statusService.Query(req.SubmissionID, true)

	e.Pong(c, eStatus, status)
	return
}
