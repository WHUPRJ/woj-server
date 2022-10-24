package problem

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/global"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/WHUPRJ/woj-server/internal/service/problem"
	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	Pid       uint   `form:"pid"`
	Title     string `form:"title"      binding:"required"`
	Statement string `form:"statement"  binding:"required"`
	IsEnabled bool   `form:"is_enabled"`
}

// Update
// @Summary     create or update a problem
// @Description create or update a problem
// @Accept      application/x-www-form-urlencoded
// @Produce     json
// @Param       pid          formData int    false "problem id, 0 for create"
// @Param       title        formData string true  "title"
// @Param       statement    formData string true  "statement"
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

	if req.Pid == 0 {
		createData := &problem.CreateData{
			Title:      req.Title,
			Statement:  req.Statement,
			ProviderID: uid,
			IsEnabled:  false,
		}
		p, status := h.problemService.Create(createData)
		e.Pong(c, status, p)
		return
	} else {
		p, status := h.problemService.Query(req.Pid, true, false)
		if status != e.Success {
			e.Pong(c, status, nil)
			return
		}
		if p.ProviderID != uid {
			e.Pong(c, e.UserUnauthorized, nil)
			return
		}

		p.Title = req.Title
		p.Statement = req.Statement
		p.IsEnabled = req.IsEnabled

		p, status = h.problemService.Update(p)
		e.Pong(c, status, p)
		return
	}
}
