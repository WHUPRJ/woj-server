package status

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/gin-gonic/gin"
)

type queryByVersionRequest struct {
	ProblemVersionID uint `form:"pvid"`
	Offset           int  `form:"offset"`
	Limit            int  `form:"limit"`
}

func (h *handler) QueryByProblemVersion(c *gin.Context) {

	claim, exist := c.Get("claim")
	if !exist {
		e.Pong(c, e.UserUnauthenticated, nil)
		return
	}

	role := claim.(*global.Claim).Role

	req := new(queryByVersionRequest)

	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, nil)
		return
	}

	if role < model.RoleAdmin {
		e.Pong(c, e.UserUnauthorized, nil)
		return
	}

	statuses, eStatus := h.statusService.QueryByVersion(req.ProblemVersionID, req.Offset, req.Limit)

	e.Pong(c, eStatus, statuses)
	return
}
