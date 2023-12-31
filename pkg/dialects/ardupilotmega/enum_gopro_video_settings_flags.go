//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package ardupilotmega

import (
	"fmt"
	"strings"
)

type GOPRO_VIDEO_SETTINGS_FLAGS uint32

const (
	// 0=NTSC, 1=PAL.
	GOPRO_VIDEO_SETTINGS_TV_MODE GOPRO_VIDEO_SETTINGS_FLAGS = 1
)

var labels_GOPRO_VIDEO_SETTINGS_FLAGS = map[GOPRO_VIDEO_SETTINGS_FLAGS]string{
	GOPRO_VIDEO_SETTINGS_TV_MODE: "GOPRO_VIDEO_SETTINGS_TV_MODE",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GOPRO_VIDEO_SETTINGS_FLAGS) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_GOPRO_VIDEO_SETTINGS_FLAGS {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GOPRO_VIDEO_SETTINGS_FLAGS) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask GOPRO_VIDEO_SETTINGS_FLAGS
	for _, label := range labels {
		found := false
		for value, l := range labels_GOPRO_VIDEO_SETTINGS_FLAGS {
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
func (e GOPRO_VIDEO_SETTINGS_FLAGS) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
