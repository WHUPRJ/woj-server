package submission

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/WHUPRJ/woj-server/internal/service/submission"
	"github.com/gin-gonic/gin"
)

type createRequest struct {
	Pid      uint   `form:"pid"       binding:"required"`
	Language string `form:"language"  binding:"required"`
	Code     string `form:"code"      binding:"required"`
}

// Create
// @Summary     create a submission
// @Description create a submission
// @Accept      application/x-www-form-urlencoded
// @Produce     json
// @Param       pid          formData int    true  "problem id"
// @Param       language     formData string true  "language"
// @Param       code         formData string true  "code"
// @Response    200 {object} e.Response ""
// @Security    Authentication
// @Router      /v1/submission/create [post]
func (h *handler) Create(c *gin.Context) {
	claim, exist := c.Get("claim")
	if !exist {
		e.Pong(c, e.UserUnauthenticated, nil)
		return
	}

	uid := claim.(*global.Claim).UID

	role := claim.(*global.Claim).Role
	req := new(createRequest)

	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, err.Error())
		return
	}

	// guest can not submit
	if role < model.RoleGeneral {
		e.Pong(c, e.UserUnauthorized, nil)
		return
	}

	createData := &submission.CreateData{
		ProblemID: req.Pid,
		UserID:    uid,
		Language:  req.Language,
		Code:      req.Code,
	}
	s, status := h.submissionService.Create(createData)
	if status != e.Success {
		e.Pong(c, status, nil)
		return
	}

	pv, status := h.problemService.QueryLatestVersion(req.Pid)
	if status != e.Success {
		e.Pong(c, status, nil)
		return
	}

	payload := &model.SubmitJudgePayload{
		ProblemVersionID: pv.ID,
		StorageKey:       pv.StorageKey,
		Submission:       *s,
	}
	_, status = h.taskService.SubmitJudge(payload)

	e.Pong(c, status, nil)
}
