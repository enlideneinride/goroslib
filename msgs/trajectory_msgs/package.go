// Package trajectory_msgs contains message definitions (autogenerated).
package trajectory_msgs

import (
	"github.com/aler9/goroslib/msg"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
	"time"
)

type JointTrajectory struct {
	msg.Package `ros:"trajectory_msgs"`
	Header      std_msgs.Header
	JointNames  []string
	Points      []JointTrajectoryPoint
}

type JointTrajectoryPoint struct {
	msg.Package   `ros:"trajectory_msgs"`
	Positions     []float64
	Velocities    []float64
	Accelerations []float64
	Effort        []float64
	TimeFromStart time.Duration
}

type MultiDOFJointTrajectory struct {
	msg.Package `ros:"trajectory_msgs"`
	Header      std_msgs.Header
	JointNames  []string
	Points      []MultiDOFJointTrajectoryPoint
}

type MultiDOFJointTrajectoryPoint struct {
	msg.Package   `ros:"trajectory_msgs"`
	Transforms    []geometry_msgs.Transform
	Velocities    []geometry_msgs.Twist
	Accelerations []geometry_msgs.Twist
	TimeFromStart time.Duration
}