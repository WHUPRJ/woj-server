package status

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/gin-gonic/gin"
)

type queryRequest struct {
	SubmissionID uint `form:"sid" binding:"required"`
}

// Query
// @Summary     query submissions by via submission id
// @Description query submissions by via submission id
// @Accept      application/x-www-form-urlencoded
// @Produce     json
// @Param       sid formData uint true "submission id"
// @Response    200 {object} e.Response "model.status"
// @Router      /v1/status/query [post]
func (h *handler) Query(c *gin.Context) {
	req := new(queryRequest)

	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, err.Error())
		return
	}

	status, eStatus := h.statusService.Query(req.SubmissionID, true)

	e.Pong(c, eStatus, status)
}
