package submission

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/gin-gonic/gin"
)

type rejudgeRequest struct {
	Sid uint `form:"sid" binding:"required"`
}

// Rejudge
// @Summary     rejudge a submission
// @Description rejudge a submission
// @Accept      application/x-www-form-urlencoded
// @Produce     json
// @Param       sid          formData int    true  "submission id"
// @Response    200 {object} e.Response ""
// @Security    Authentication
// @Router      /v1/submission/rejudge [post]
func (h *handler) Rejudge(c *gin.Context) {
	claim, exist := c.Get("claim")
	if !exist {
		e.Pong(c, e.UserUnauthenticated, nil)
		return
	}

	role := claim.(*global.Claim).Role
	req := new(rejudgeRequest)

	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, err.Error())
		return
	}

	// only admin can rejudge
	if role < model.RoleAdmin {
		e.Pong(c, e.UserUnauthorized, nil)
		return
	}

	s, status := h.submissionService.QueryBySid(req.Sid, false)
	if status != e.Success {
		e.Pong(c, status, nil)
		return
	}

	pv, status := h.problemService.QueryLatestVersion(s.ProblemID)
	if status != e.Success {
		e.Pong(c, status, nil)
		return
	}

	_, status = h.taskService.SubmitJudge(&model.SubmitJudgePayload{
		ProblemVersionID: pv.ID,
		StorageKey:       pv.StorageKey,
		Submission:       *s,
	})

	e.Pong(c, status, nil)
}
