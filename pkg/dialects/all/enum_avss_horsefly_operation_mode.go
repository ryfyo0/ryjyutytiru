//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package all

import (
	"errors"
)

type AVSS_HORSEFLY_OPERATION_MODE uint32

const (
	// In manual control mode
	MODE_HORSEFLY_MANUAL_CTRL AVSS_HORSEFLY_OPERATION_MODE = 0
	// In auto takeoff mode
	MODE_HORSEFLY_AUTO_TAKEOFF AVSS_HORSEFLY_OPERATION_MODE = 1
	// In auto landing mode
	MODE_HORSEFLY_AUTO_LANDING AVSS_HORSEFLY_OPERATION_MODE = 2
	// In go home mode
	MODE_HORSEFLY_NAVI_GO_HOME AVSS_HORSEFLY_OPERATION_MODE = 3
	// In drop mode
	MODE_HORSEFLY_DROP AVSS_HORSEFLY_OPERATION_MODE = 4
)

var labels_AVSS_HORSEFLY_OPERATION_MODE = map[AVSS_HORSEFLY_OPERATION_MODE]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e AVSS_HORSEFLY_OPERATION_MODE) MarshalText() ([]byte, error) {
	if l, ok := labels_AVSS_HORSEFLY_OPERATION_MODE[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_AVSS_HORSEFLY_OPERATION_MODE = map[string]AVSS_HORSEFLY_OPERATION_MODE{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *AVSS_HORSEFLY_OPERATION_MODE) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_AVSS_HORSEFLY_OPERATION_MODE[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e AVSS_HORSEFLY_OPERATION_MODE) String() string {
	if l, ok := labels_AVSS_HORSEFLY_OPERATION_MODE[e]; ok {
		return l
	}
	return "invalid value"
}