//nolint:golint
package vision_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
	"github.com/aler9/goroslib/pkg/msgs/std_msgs"
)

type VisionInfo struct {
	msg.Package      `ros:"vision_msgs"`
	Header           std_msgs.Header
	Method           string
	DatabaseLocation string
	DatabaseVersion  int32
}
