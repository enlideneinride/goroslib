package protoudp

import (
	"encoding/binary"
	"fmt"
)

const (
	maxPayloadSize = 1492
)

// Opcode is the opcode of a Frame.
type Opcode uint8

// Standard opcodes.
const (
	Data0 Opcode = 0
	DataN Opcode = 1
	Ping  Opcode = 2
	Error Opcode = 3
)

// Frame is a UDPROS frame.
type Frame struct {
	ConnectionID uint32
	Opcode       Opcode
	MessageID    uint8
	BlockID      uint16
	Payload      []byte
}

func (f *Frame) decode(byts []byte) error {
	if len(byts) < 8 {
		return fmt.Errorf("invalid length")
	}

	f.ConnectionID = binary.LittleEndian.Uint32(byts[0:4])
	f.Opcode = Opcode(byts[4])
	f.MessageID = byts[5]
	f.BlockID = binary.LittleEndian.Uint16(byts[6:8])
	f.Payload = byts[8:]

	return nil
}

func (f *Frame) encode() ([]byte, error) {
	byts := make([]byte, 8+len(f.Payload))

	binary.LittleEndian.PutUint32(byts[:4], f.ConnectionID)
	byts[4] = uint8(f.Opcode)
	byts[5] = f.MessageID
	binary.LittleEndian.PutUint16(byts[6:8], f.BlockID)
	copy(byts[8:], f.Payload)

	return byts, nil
}

// FramesForPayload generates frames for the given payload.
func FramesForPayload(connID uint32, messageID uint8, payload []byte) []*Frame {
	var ret []*Frame
	lbyts := len(payload)

	for i := 0; i < lbyts; i += maxPayloadSize {
		f := &Frame{
			ConnectionID: connID,
			Opcode: func() Opcode {
				if i == 0 {
					return Data0
				}
				return DataN
			}(),
			MessageID: messageID,
			BlockID: func() uint16 {
				// return block count
				if i == 0 {
					if (lbyts % maxPayloadSize) == 0 {
						return uint16(lbyts / maxPayloadSize)
					}
					return uint16((lbyts / maxPayloadSize) + 1)
				}

				// return current block id
				return uint16(i / maxPayloadSize)
			}(),
			Payload: func() []byte {
				j := i + maxPayloadSize
				if j > lbyts {
					j = lbyts
				}
				return payload[i:j]
			}(),
		}
		ret = append(ret, f)
	}

	return ret
}
