//nolint:golint
package mavros_msgs

import (
	"time"

	"github.com/aler9/goroslib/pkg/msg"
	"github.com/aler9/goroslib/pkg/msgs/std_msgs"
)

type LogEntry struct {
	msg.Package `ros:"mavros_msgs"`
	Header      std_msgs.Header
	Id          uint16
	NumLogs     uint16
	LastLogNum  uint16
	TimeUtc     time.Time
	Size        uint32
}