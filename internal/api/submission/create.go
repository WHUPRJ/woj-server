package submission

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/service/submission"
	"github.com/gin-gonic/gin"
)

type createRequest struct {
	Pid      uint   `form:"pid"       binding:"required"`
	Language string `form:"language"  binding:"required"`
	Code     string `form:"statement" binding:"required"`
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

	req := new(createRequest)

	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, err.Error())
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

	_, status = h.taskService.SubmitJudge(pv.ID, pv.StorageKey, *s)
	e.Pong(c, status, nil)
	return
}
