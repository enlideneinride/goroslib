//nolint:golint
package mavros_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
	"github.com/aler9/goroslib/pkg/msgs/geometry_msgs"
	"github.com/aler9/goroslib/pkg/msgs/std_msgs"
)

type HilSensor struct {
	msg.Package   `ros:"mavros_msgs"`
	Header        std_msgs.Header
	Acc           geometry_msgs.Vector3
	Gyro          geometry_msgs.Vector3
	Mag           geometry_msgs.Vector3
	AbsPressure   float32
	DiffPressure  float32
	PressureAlt   float32
	Temperature   float32
	FieldsUpdated uint32
}
