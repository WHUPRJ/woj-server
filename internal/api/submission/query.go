package submission

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/internal/model"
	"github.com/gin-gonic/gin"
)

type queryRequest struct {
	Pid    uint `form:"pid"`
	Uid    uint `form:"uid"`
	Offset int  `form:"offset"`
	Limit  int  `form:"limit"`
}

type queryResponse struct {
	Submission model.Submission `json:"submission"`
	Point      int32            `json:"point"`
}

// Query
// @Summary     Query submissions
// @Description Query submissions
// @Accept      application/x-www-form-urlencoded
// @Produce     json
// @Param       pid formData uint false "problem id"
// @Param       uid formData uint false "user id"
// @Param       offset formData int false "start position"
// @Param       limit formData int false "limit number of records"
// @Response    200 {object} e.Response "queryResponse"
// @Router      /v1/submission/query [post]

func (h *handler) Query(c *gin.Context) {
	req := new(queryRequest)

	if err := c.ShouldBind(req); err != nil {
		e.Pong(c, e.InvalidParameter, err.Error())
		return
	}

	if req.Pid == 0 && req.Uid == 0 {
		e.Pong(c, e.InvalidParameter, nil)
		return
	}

	submissions, status := h.submissionService.Query(req.Pid, req.Uid, req.Offset, req.Limit)

	var response []*queryResponse

	for _, submission := range submissions {
		currentStatus, _ := h.statusService.Query(submission.ID, false)
		var currentPoint int32

		if currentStatus == nil {
			currentPoint = -1
		} else {
			currentPoint = currentStatus.Point
		}

		newResponse := &queryResponse{
			Submission: *submission,
			Point:      currentPoint,
		}

		newResponse.Submission.Code = ""

		response = append(response, newResponse)
	}

	e.Pong(c, status, submissions)
	return

}
