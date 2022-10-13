package problem

import (
	"github.com/WHUPRJ/woj-server/global"
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/model"
	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	Pid         uint   `form:"pid"`
	Title       string `form:"title"        binding:"required"`
	Content     string `form:"content"      binding:"required"`
	TimeLimit   uint   `form:"time_limit"   binding:"required"`
	MemoryLimit uint   `form:"memory_limit" binding:"required"`
	IsEnabled   bool   `form:"is_enabled"`
}

// Update
// @Summary     create or update a problem
// @Description create or update a problem
// @Accept      application/x-www-form-urlencoded
// @Produce     json
// @Param       pid          formData int    false "problem id, 0 for create"
// @Param       title        formData string true  "title"
// @Param       content      formData string true  "content"
// @Param       time_limit   formData int    true  "time limit in ms"
// @Param       memory_limit formData int    true  "memory limit in kb"
// @Param       is_enabled   formData bool   false "is enabled"
// @Response    200 {object} e.Response "problem info without provider information"
// @Security    Authentication
// @Router      /v1/problem/update [post]
func (h *handler) Update(c *gin.Context) {
	claim, exist := c.Get("claim")
	if !exist {
		e.Pong(c, e.UserUnauthenticated, nil)
		return
	}

	uid := claim.(*global.Claim).UID
	role := claim.(*global.Claim).Role
	if role < model.RoleAdmin {
		e.Pong(c, e.UserUnauthorized, nil)
		return
	}

	req := new(updateRequest)
	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, err.Error())
		return
	}

	problem := &model.Problem{
		Title:       req.Title,
		Content:     req.Content,
		TimeLimit:   req.TimeLimit,
		MemoryLimit: req.MemoryLimit,
		IsEnabled:   req.IsEnabled,
	}

	if req.Pid == 0 {
		problem, status := h.problemService.Create(uid, problem)
		e.Pong(c, status, problem)
		return
	} else {
		inDb, status := h.problemService.Query(req.Pid)
		if status != e.Success && status != e.ProblemNotAvailable {
			e.Pong(c, status, nil)
			return
		}

		if inDb.ProviderID != uid {
			e.Pong(c, e.UserUnauthorized, nil)
			return
		}

		problem, status := h.problemService.Update(req.Pid, problem)
		e.Pong(c, status, problem)
		return
	}
}
