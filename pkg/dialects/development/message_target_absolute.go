//autogenerated:yes
//nolint:revive,misspell,govet,lll
package development

// Current motion information from sensors on a target
type MessageTargetAbsolute struct {
	// Timestamp (UNIX epoch time).
	Timestamp uint64
	// The ID of the target if multiple targets are present
	Id uint8
	// Bitmap to indicate the sensor's reporting capabilities
	SensorCapabilities TARGET_ABSOLUTE_SENSOR_CAPABILITY_FLAGS `mavenum:"uint8"`
	// Target's latitude (WGS84)
	Lat int32
	// Target's longitude (WGS84)
	Lon int32
	// Target's altitude (AMSL)
	Alt float32
	// Target's velocity in its body frame
	Vel [3]float32
	// Linear target's acceleration in its body frame
	Acc [3]float32
	// Quaternion of the target's orientation from its body frame to the vehicle's NED frame.
	QTarget [4]float32
	// Target's roll, pitch and yaw rates
	Rates [3]float32
	// Standard deviation of horizontal (eph) and vertical (epv) position errors
	PositionStd [2]float32
	// Standard deviation of the target's velocity in its body frame
	VelStd [3]float32
	// Standard deviation of the target's acceleration in its body frame
	AccStd [3]float32
}

// GetID implements the message.Message interface.
func (*MessageTargetAbsolute) GetID() uint32 {
	return 510
}
