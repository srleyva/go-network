package main

import (
	"log"

	"github.com/songgao/water"
	"github.com/songgao/water/waterutil"

	"github.com/srleyva/tcp-go/pkg/icmp"
	"github.com/srleyva/tcp-go/pkg/ipv4"
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
		bytesRead, err := nic.Read(buffer)
		if err != nil {
			panic(err)
		}

		if waterutil.IsIPv4(buffer) {

			proto := buffer[9] // Octet 9 or bit 72 in the ip packet
			ipv4Packet := ipv4.Packet{
				SrcIP:    waterutil.IPv4Source(buffer),
				SrcPort:  waterutil.IPv4SourcePort(buffer),
				DestIP:   waterutil.IPv4Destination(buffer),
				DestPort: waterutil.IPv4DestinationPort(buffer),
			}

			switch proto {
			case waterutil.ICMP:
				// TODO ICMP
				log.Printf("Received ICMP from %s:%d responding that I am alive and kicking", ipv4Packet.SrcIP, ipv4Packet.SrcPort)

				payload := waterutil.IPv4Payload(buffer)
				ipv4Packet.Payload, err = icmp.NewICMP(payload, bytesRead)
				if err != nil {
					log.Printf("err marshalling payload into ICMP: %s", err)
					continue
				}

				respPayload, err := ipv4Packet.Payload.HandlePacket()
				if err != nil {
					log.Printf("err handling packet for ICMP: %s", err)
				}

				resp := ipv4.Packet{
					Proto:    waterutil.ICMP,
					SrcIP:    ipv4Packet.DestIP,
					SrcPort:  ipv4Packet.DestPort,
					DestIP:   ipv4Packet.SrcIP,
					DestPort: ipv4Packet.SrcPort,
					Payload:  respPayload,
				}

				response := resp.ToByteArray(buffer)

				_, err = nic.Write(response)
				if err != nil {
					log.Printf("err sending packet: %s", err)
				}
				log.Print("My Response was written successfully")

			case waterutil.TCP:
				log.Printf(
					"%s:%d -> %s:%d proto: TCP",
					ipv4Packet.SrcIP,
					ipv4Packet.SrcPort,
					ipv4Packet.DestIP,
					ipv4Packet.DestPort)
			case waterutil.UDP:
				// TODO UDP
				panic("Not implemented")

			}

		} else {
			log.Print("Weird not ipv4 packet")
		}
	}

}
