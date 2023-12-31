//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package avssuas

import (
	"fmt"
	"strings"
)

type MAV_AVSS_COMMAND_FAILURE_REASON uint32

const (
	// AVSS defined command failure reason. PRS not steady.
	PRS_NOT_STEADY MAV_AVSS_COMMAND_FAILURE_REASON = 1
	// AVSS defined command failure reason. PRS DTM not armed.
	PRS_DTM_NOT_ARMED MAV_AVSS_COMMAND_FAILURE_REASON = 2
	// AVSS defined command failure reason. PRS OTM not armed.
	PRS_OTM_NOT_ARMED MAV_AVSS_COMMAND_FAILURE_REASON = 3
)

var labels_MAV_AVSS_COMMAND_FAILURE_REASON = map[MAV_AVSS_COMMAND_FAILURE_REASON]string{
	PRS_NOT_STEADY:    "PRS_NOT_STEADY",
	PRS_DTM_NOT_ARMED: "PRS_DTM_NOT_ARMED",
	PRS_OTM_NOT_ARMED: "PRS_OTM_NOT_ARMED",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e MAV_AVSS_COMMAND_FAILURE_REASON) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_MAV_AVSS_COMMAND_FAILURE_REASON {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *MAV_AVSS_COMMAND_FAILURE_REASON) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask MAV_AVSS_COMMAND_FAILURE_REASON
	for _, label := range labels {
		found := false
		for value, l := range labels_MAV_AVSS_COMMAND_FAILURE_REASON {
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
func (e MAV_AVSS_COMMAND_FAILURE_REASON) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
