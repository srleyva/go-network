package main

import (
	"github.com/songgao/water"
	"github.com/songgao/water/waterutil"
	"log"
)

func main() {
	nic, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Interface Name: %s\n", nic.Name())

	buffer := make([]byte, 1504)

	for {
		_, err := nic.Read(buffer)
		if err != nil {
			panic(err)
		}

		if waterutil.IsIPv4(buffer) {
			proto := buffer[9] // Octet 9 or bit 72 in the ip packet
			src := waterutil.IPv4Source(buffer)
			srcPort := waterutil.IPv4SourcePort(buffer)
			dst := waterutil.IPv4Destination(buffer)
			dstPort := waterutil.IPv4DestinationPort(buffer)

			switch proto {
			case waterutil.ICMP:
				// TODO ICMP
				log.Printf("%s:%d -> %s:%d proto: ICMP", src, srcPort, dst, dstPort)
			case waterutil.TCP:
				log.Printf("%s:%d -> %s:%d proto: TCP", src, srcPort, dst, dstPort)
			case waterutil.UDP:
				// TODO UDP
				log.Printf("%s:%d -> %s:%d proto: UDP", src, srcPort, dst, dstPort)

			}

		} else {
			log.Print("Weird not ipv4 packet")
		}
	}

}
