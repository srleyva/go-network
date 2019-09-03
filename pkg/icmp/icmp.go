package icmp

import (
	"encoding/binary"

	"github.com/srleyva/tcp-go/pkg/ipv4"
)

const (
	EchoReply   = 0
	EchoRequest = 8
)

const (
	CheckSumsDoNotMatch = "Checksums do not match"
)

// ICMP fullfills the payload interface for ICMP packets
type ICMP struct {
	Type     uint8
	Code     uint8
	Checksum uint16
	Data     []byte
}

// NewICMP takes raw bytes and marshals them into ICMP
func NewICMP(buffer []byte, bytesRead int) (*ICMP, error) {
	checksum := binary.BigEndian.Uint16([]byte{buffer[2], buffer[3]})

	icmpPacket := &ICMP{
		Type:     uint8(buffer[0]),
		Code:     uint8(buffer[1]),
		Checksum: checksum,
		Data:     buffer[4:bytesRead],
	}

	// Validate using checksum
	// calculatedChecksum := icmpPacket.calculateChecksum()

	// if calculatedChecksum != checksum {
	// 	log.Printf("Calculated Checksum: %d Expected Checksum: %d", calculatedChecksum, checksum)
	// 	return nil, fmt.Errorf(CheckSumsDoNotMatch)
	// }

	return icmpPacket, nil
}

// HandlePacket Create a Echo Reply to response to Request
func (i *ICMP) HandlePacket() (ipv4.Payload, error) {
	var response ICMP
	switch i.Type {
	case EchoRequest:
		identifier := uint16(0)
		SeqNumber := uint16(0)
		dataBytes := make([]byte, 4) // Allow for Identifier and Sequence Number plus the data recieved in request

		binary.BigEndian.PutUint16(dataBytes[0:2], uint16(identifier))
		binary.BigEndian.PutUint16(dataBytes[2:4], uint16(SeqNumber))

		dataBytes = append(dataBytes, i.Data...)

		response = ICMP{
			Type: EchoReply,
			Code: 0,
			Data: dataBytes,
		}

		response.Checksum = response.calculateChecksum()
	}

	return &response, nil
}

// ToByteArray returns a []byte representation of the struct
func (i *ICMP) ToByteArray() ([]byte, error) {

	returnValue := []byte{byte(i.Type), byte(i.Code)}

	checkSum := make([]byte, 2)
	binary.BigEndian.PutUint16(checkSum, i.Checksum)

	returnValue = append(returnValue, checkSum...)
	returnValue = append(returnValue, i.Data...)

	return returnValue, nil
}

// TODO fix
func (i *ICMP) calculateChecksum() uint16 {
	b := []byte{byte(i.Type), byte(i.Code), 0, 0}
	b = append(b, i.Data...)

	csumcv := len(b) - 1 // checksum coverage
	s := uint32(0)
	for i := 0; i < csumcv; i += 2 {
		s += uint32(b[i+1])<<8 | uint32(b[i])
	}
	if csumcv&1 == 0 {
		s += uint32(b[csumcv])
	}
	s = s>>16 + s&0xffff
	s = s + s>>16
	return ^uint16(s)
}
