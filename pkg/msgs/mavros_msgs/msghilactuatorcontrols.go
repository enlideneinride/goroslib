//nolint:golint
package mavros_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
	"github.com/aler9/goroslib/pkg/msgs/std_msgs"
)

type HilActuatorControls struct {
	msg.Package `ros:"mavros_msgs"`
	Header      std_msgs.Header
	Controls    [16]float32
	Mode        uint8
	Flags       uint64
}
