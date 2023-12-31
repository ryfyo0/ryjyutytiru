//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"fmt"
	"strings"
)

type PID_TUNING_AXIS uint32

const (
	PID_TUNING_ROLL    PID_TUNING_AXIS = 1
	PID_TUNING_PITCH   PID_TUNING_AXIS = 2
	PID_TUNING_YAW     PID_TUNING_AXIS = 3
	PID_TUNING_ACCZ    PID_TUNING_AXIS = 4
	PID_TUNING_STEER   PID_TUNING_AXIS = 5
	PID_TUNING_LANDING PID_TUNING_AXIS = 6
)

var labels_PID_TUNING_AXIS = map[PID_TUNING_AXIS]string{
	PID_TUNING_ROLL:    "PID_TUNING_ROLL",
	PID_TUNING_PITCH:   "PID_TUNING_PITCH",
	PID_TUNING_YAW:     "PID_TUNING_YAW",
	PID_TUNING_ACCZ:    "PID_TUNING_ACCZ",
	PID_TUNING_STEER:   "PID_TUNING_STEER",
	PID_TUNING_LANDING: "PID_TUNING_LANDING",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e PID_TUNING_AXIS) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_PID_TUNING_AXIS {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *PID_TUNING_AXIS) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask PID_TUNING_AXIS
	for _, label := range labels {
		found := false
		for value, l := range labels_PID_TUNING_AXIS {
			if l == label {
				mask |= value
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("invalid label '%s'", label)
		}
	}
	*e = mask
	return nil
}

// String implements the fmt.Stringer interface.
func (e PID_TUNING_AXIS) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
