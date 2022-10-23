package status

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/gin-gonic/gin"
)

type queryByVersionRequest struct {
	ProblemVersionID uint `form:"pvid"   binding:"required"`
	Offset           int  `form:"offset"`
	Limit            int  `form:"limit"  binding:"required"`
}

// QueryByProblemVersion
// @Summary     query submissions by problem version (admin only)
// @Description query submissions by problem version (admin only)
// @Accept      application/x-www-form-urlencoded
// @Produce     json
// @Param       pvid formData uint true "problem version id"
// @Param       offset formData int true "start position"
// @Param       limit formData int true "limit number of records"
// @Response    200 {object} e.Response "[]*model.status"
// @Router      /v1/status/query/problem_version [post]
func (h *handler) QueryByProblemVersion(c *gin.Context) {

	claim, exist := c.Get("claim")
	if !exist {
		e.Pong(c, e.UserUnauthenticated, nil)
		return
	}

	role := claim.(*global.Claim).Role

	req := new(queryByVersionRequest)

	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, err.Error())
		return
	}

	if role < model.RoleAdmin {
		e.Pong(c, e.UserUnauthorized, nil)
		return
	}

	statuses, eStatus := h.statusService.QueryByVersion(req.ProblemVersionID, req.Offset, req.Limit)

	e.Pong(c, eStatus, statuses)
}
