//autogenerated:yes
//nolint:revive,misspell,govet,lll
package csairlink

// Authorization package
type MessageAirlinkAuth struct {
	// Login
	Login string `mavlen:"50"`
	// Password
	Password string `mavlen:"50"`
}

// GetID implements the message.Message interface.
func (*MessageAirlinkAuth) GetID() uint32 {
	return 52000
}
