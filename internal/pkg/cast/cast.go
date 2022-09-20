package cast

import (
	"fmt"
	"github.com/WHUPRJ/woj-server/internal/e"
	"strconv"
)

// ToString Only supports some primitives and internal types
func ToString(obj interface{}) string {
	switch t := obj.(type) {
	case bool:
		return strconv.FormatBool(t)
	case []byte:
		return string(t)
	case e.Status:
		return t.String()
	case error:
		return t.Error()
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	case int:
		return strconv.Itoa(t)
	case string:
		return t
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", obj)
	}
}
