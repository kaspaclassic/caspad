package appmessage

// MsgDonePruningPointUTXOSetChunks represents a caspa DonePruningPointUTXOSetChunks message
type MsgDonePruningPointUTXOSetChunks struct {
	baseMessage
}

// Command returns the protocol command string for the message
func (msg *MsgDonePruningPointUTXOSetChunks) Command() MessageCommand {
	return CmdDonePruningPointUTXOSetChunks
}

// NewMsgDonePruningPointUTXOSetChunks returns a new MsgDonePruningPointUTXOSetChunks.
func NewMsgDonePruningPointUTXOSetChunks() *MsgDonePruningPointUTXOSetChunks {
	return &MsgDonePruningPointUTXOSetChunks{}
}
