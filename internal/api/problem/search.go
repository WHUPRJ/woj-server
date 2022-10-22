package problem

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/gin-gonic/gin"
)

type searchRequest struct {
	Search string `form:"search"`
}

// Search
// @Summary     get detail of a problem
// @Description get detail of a problem
// @Accept      application/x-www-form-urlencoded
// @Produce     json
// @Param       search formData string false "word search"
// @Response    200 {object} e.Response "problemset"
// @Router      /v1/problem/search [post]
func (h *handler) Search(c *gin.Context) {
	req := new(searchRequest)

	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, err.Error())
		return
	}

	// TODO: pagination
	if req.Search == "" {
		// TODO: query without LIKE
		problem, status := h.problemService.QueryFuzz(req.Search, true, true)
		e.Pong(c, status, problem)
		return
	} else {
		problem, status := h.problemService.QueryFuzz(req.Search, true, true)
		e.Pong(c, status, problem)
		return
	}
}
