// Copyright (c) 2014-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"io"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

// MsgBFD implements the Message interface and represents a bitcoin
// BFD message which is used to submit a block filter digest.
//
// This message was not added until protocol version XXX.
type MsgBFD struct {
	BlockHeight int32
	BlockHash chainhash.Hash
	Digest    MsgFilterLoad
}

// BtcDecode decodes r using the bitcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgBFD) BtcDecode(r io.Reader, pver uint32, enc MessageEncoding) error {
	// if pver < BIP0037Version {
	// 	str := fmt.Sprintf("filterload message invalid for protocol "+
	// 		"version %d", pver)
	// 	return messageError("MsgFilterLoad.BtcDecode", str)
	// }

	err := readElements(r, &msg.BlockHeight, &msg.BlockHash)
	if err != nil {
		return err
	}

	return msg.Digest.BtcDecode(r, pver, enc)
}

// BtcEncode encodes the receiver to w using the bitcoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgBFD) BtcEncode(w io.Writer, pver uint32, enc MessageEncoding) error {
	// if pver < BIP0037Version {
	// 	str := fmt.Sprintf("filterload message invalid for protocol "+
	// 		"version %d", pver)
	// 	return messageError("MsgFilterLoad.BtcEncode", str)
	// }

	err := writeElements(w, &msg.BlockHeight, &msg.BlockHash)
	if err != nil {
		return err
	}

	return msg.Digest.BtcEncode(w, pver, enc)
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgBFD) Command() string {
	return CmdBFD
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgBFD) MaxPayloadLength(pver uint32) uint32 {
	// Num filter bytes (varInt) + filter + 4 bytes hash funcs +
	// 4 bytes tweak + 1 byte flags.
	return uint32(VarIntSerializeSize(MaxFilterLoadFilterSize)) +
		MaxFilterLoadFilterSize + 9 + 36
}

// NewMsgBFD returns a new bitcoin bfd message that conforms to
// the Message interface.  See MsgBFD for details.
func NewMsgBFD(blockHeight int32, blockHash chainhash.Hash, filter []byte, hashFuncs uint32, tweak uint32, flags BloomUpdateType) *MsgBFD {
	return &MsgBFD{
		BlockHeight: blockHeight,
		BlockHash:   blockHash,
		Digest:      MsgFilterLoad{
			Filter:      filter,
			HashFuncs:   hashFuncs,
			Tweak:       tweak,
			Flags:       flags,
		},
	}
}
