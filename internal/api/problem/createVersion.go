package problem

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/WHUPRJ/woj-server/internal/service/problem"
	"github.com/gin-gonic/gin"
)

type createVersionRequest struct {
	ProblemID  uint   `form:"pid"         binding:"required"`
	StorageKey string `form:"storage_key" binding:"required"`
}

// CreateVersion
// @Summary     create a problem version
// @Description create a problem version
// @Accept      application/x-www-form-urlencoded
// @Produce     json
// @Param       pid         formData int    true  "problem id"
// @Param       storage_key formData string true  "storage key"
// @Response    200 {object} e.Response ""
// @Security    Authentication
// @Router      /v1/problem/create_version [post]
func (h *handler) CreateVersion(c *gin.Context) {
	claim, exist := c.Get("claim")
	if !exist {
		e.Pong(c, e.UserUnauthenticated, nil)
		return
	}

	// uid := claim.(*global.Claim).UID

	role := claim.(*global.Claim).Role
	req := new(createVersionRequest)

	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, err.Error())
		return
	}

	// guest can not submit
	if role < model.RoleAdmin {
		e.Pong(c, e.UserUnauthorized, nil)
		return
	}

	// TODO: check pid exist

	createVersionData := &problem.CreateVersionData{
		ProblemID:  req.ProblemID,
		StorageKey: req.StorageKey,
	}
	pv, status := h.problemService.CreateVersion(createVersionData)
	if status != e.Success {
		e.Pong(c, status, nil)
		return
	}

	payload := &model.ProblemBuildPayload{
		ProblemVersionID: pv.ID,
		StorageKey:       pv.StorageKey,
	}
	_, status = h.taskService.ProblemBuild(payload)
	e.Pong(c, status, nil)
}
