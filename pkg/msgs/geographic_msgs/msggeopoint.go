//nolint:golint
package geographic_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
)

type GeoPoint struct {
	msg.Package `ros:"geographic_msgs"`
	Latitude    float64
	Longitude   float64
	Altitude    float64
}
