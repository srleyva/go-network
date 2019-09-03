package ipv4

import (
	"net"
)

// Packet represents the field in an IP Packet
type Packet struct {
	SrcIP    net.IP
	SrcPort  uint16
	DestIP   net.IP
	DestPort uint16
	Proto    int
	Payload  Payload
}

// ToByteArray serializes Packet to Byte Array
func (p *Packet) ToByteArray(buffer []byte) []byte {

	// ATM this is not insanely valid as checksum calc needs to happen
	// but since were just swapping the IP src and dest the checksum is valid
	dest := []byte{buffer[12], buffer[13], buffer[14], buffer[15]}
	src := []byte{buffer[16], buffer[17], buffer[18], buffer[19]}

	// Swap Src and Dest
	byteLocation := 12
	for i := 0; i < len(dest); i++ {
		buffer[byteLocation] = src[i]
		buffer[byteLocation+4] = dest[i]
		byteLocation++
	}

	return buffer

}

// Payload represents the payload in an IPv4 packet
// This could be TCP, ICMP, UDP, etc
type Payload interface {
	HandlePacket() (Payload, error)
	ToByteArray() ([]byte, error)
}
