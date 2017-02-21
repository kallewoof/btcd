// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"io"
)

// MsgGetBFD implements the Message interface and represents a bitcoin
// getbfd message.  It is used to request the latest block filter digest.
type MsgGetBFD struct {
	BlockHeight int32
}

// BtcDecode decodes r using the bitcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgGetBFD) BtcDecode(r io.Reader, pver uint32, enc MessageEncoding) error {
	err := readElements(r, &msg.BlockHeight)
	if err != nil {
		return err
	}

	return nil
}

// BtcEncode encodes the receiver to w using the bitcoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgGetBFD) BtcEncode(w io.Writer, pver uint32, enc MessageEncoding) error {
	err := writeElements(w, &msg.BlockHeight)
	if err != nil {
		return err
	}

	return nil
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgGetBFD) Command() string {
	return CmdGetBFD
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgGetBFD) MaxPayloadLength(pver uint32) uint32 {
	// block height 4 bytes
	return 4
}

// NewMsgGetBFD returns a new bitcoin getblocks message that conforms to the
// Message interface using the passed parameters and defaults for the remaining
// fields.
func NewMsgGetBFD(blockHeight int32) *MsgGetBFD {
	return &MsgGetBFD{
		BlockHeight: blockHeight,
	}
}
