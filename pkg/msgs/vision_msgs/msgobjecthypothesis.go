//nolint:golint
package vision_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
)

type ObjectHypothesis struct {
	msg.Package `ros:"vision_msgs"`
	Id          int64
	Score       float64
}
