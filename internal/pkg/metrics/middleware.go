package metrics

import (
	"github.com/WHUPRJ/woj-server/internal/e"
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func (m *Metrics) SetLogPaths(paths []string) {
	m.logPaths = paths
}

func (m *Metrics) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.Request.URL.String()
		method := c.Request.Method

		if !utils.Contains(m.logPaths, func(path string) bool { return strings.HasPrefix(url, path) }) {
			c.Next()
			return
		}

		start := time.Now()
		c.Next()
		elapsed := float64(time.Since(start)) / float64(time.Millisecond)

		status := c.Writer.Status()
		success := !c.IsAborted() && (status == http.StatusOK)
		errCode, _ := c.Get("err")
		err, ok := errCode.(e.Status)
		if !ok {
			success = false
			err = e.Unknown
		} else if err != e.Success {
			success = false
		}

		m.Record(method, url, success, status, err, elapsed)
	}
}
