package appmessage

// MsgUnexpectedPruningPoint represents a Kalibrium UnexpectedPruningPoint message
type MsgUnexpectedPruningPoint struct {
	baseMessage
}

// Command returns the protocol command string for the message
func (msg *MsgUnexpectedPruningPoint) Command() MessageCommand {
	return CmdUnexpectedPruningPoint
}

// NewMsgUnexpectedPruningPoint returns a new Kalibrium UnexpectedPruningPoint message
func NewMsgUnexpectedPruningPoint() *MsgUnexpectedPruningPoint {
	return &MsgUnexpectedPruningPoint{}
}
