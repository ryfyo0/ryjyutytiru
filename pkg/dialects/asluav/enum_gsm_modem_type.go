//autogenerated:yes
//nolint:revive,misspell,govet,lll,dupl,gocritic
package asluav

import (
	"fmt"
	"strings"
)

type GSM_MODEM_TYPE uint32

const (
	// not specified
	GSM_MODEM_TYPE_UNKNOWN GSM_MODEM_TYPE = 0
	// HUAWEI LTE USB Stick E3372
	GSM_MODEM_TYPE_HUAWEI_E3372 GSM_MODEM_TYPE = 1
)

var labels_GSM_MODEM_TYPE = map[GSM_MODEM_TYPE]string{
	GSM_MODEM_TYPE_UNKNOWN:      "GSM_MODEM_TYPE_UNKNOWN",
	GSM_MODEM_TYPE_HUAWEI_E3372: "GSM_MODEM_TYPE_HUAWEI_E3372",
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e GSM_MODEM_TYPE) MarshalText() ([]byte, error) {
	var names []string
	for mask, label := range labels_GSM_MODEM_TYPE {
		if e&mask == mask {
			names = append(names, label)
		}
	}
	return []byte(strings.Join(names, " | ")), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (e *GSM_MODEM_TYPE) UnmarshalText(text []byte) error {
	labels := strings.Split(string(text), " | ")
	var mask GSM_MODEM_TYPE
	for _, label := range labels {
		found := false
		for value, l := range labels_GSM_MODEM_TYPE {
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
func (e GSM_MODEM_TYPE) String() string {
	val, _ := e.MarshalText()
	return string(val)
}
