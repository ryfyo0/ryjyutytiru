//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package uavionix

import (
	"fmt"
	"strings"
)

// Status for ADS-B transponder dynamic input
type UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX uint32

const (
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_NONE_0 UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX = 0
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_NONE_1 UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX = 1
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_2D     UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX = 2
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_3D     UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX = 3
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_DGPS   UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX = 4
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_RTK    UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX = 5
)

var labels_UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX = map[UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX]string{
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_NONE_0: "UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_NONE_0",
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_NONE_1: "UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_NONE_1",
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_2D:     "UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_2D",
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_3D:     "UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_3D",
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_DGPS:   "UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_DGPS",
	UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_RTK:    "UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX_RTK",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX
	for _, label := range labels {
		found := false
		for value, l := range labels_UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX {
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
func (e UAVIONIX_ADSB_OUT_DYNAMIC_GPS_FIX) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
