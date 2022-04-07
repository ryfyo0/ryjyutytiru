//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl
package all

import (
	"errors"
)

type ICAROUS_TRACK_BAND_TYPES uint32

const (
	ICAROUS_TRACK_BAND_TYPE_NONE     ICAROUS_TRACK_BAND_TYPES = 0
	ICAROUS_TRACK_BAND_TYPE_NEAR     ICAROUS_TRACK_BAND_TYPES = 1
	ICAROUS_TRACK_BAND_TYPE_RECOVERY ICAROUS_TRACK_BAND_TYPES = 2
)

var labels_ICAROUS_TRACK_BAND_TYPES = map[ICAROUS_TRACK_BAND_TYPES]string{}

// MarshalText implements the encoding.TextMarshaler interface.
func (e ICAROUS_TRACK_BAND_TYPES) MarshalText() ([]byte, error) {
	if l, ok := labels_ICAROUS_TRACK_BAND_TYPES[e]; ok {
		return []byte(l), nil
	}
	return nil, errors.New("invalid value")
}

var reverseLabels_ICAROUS_TRACK_BAND_TYPES = map[string]ICAROUS_TRACK_BAND_TYPES{}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *ICAROUS_TRACK_BAND_TYPES) UnmarshalText(text []byte) error {
	if rl, ok := reverseLabels_ICAROUS_TRACK_BAND_TYPES[string(text)]; ok {
		*e = rl
		return nil
	}
	return errors.New("invalid value")
}

// String implements the fmt.Stringer interface.
func (e ICAROUS_TRACK_BAND_TYPES) String() string {
	if l, ok := labels_ICAROUS_TRACK_BAND_TYPES[e]; ok {
		return l
	}
	return "invalid value"
}