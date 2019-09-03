module github.com/srleyva/tcp-go/pkg/icmp

go 1.12

require (
	github.com/srleyva/tcp-go/pkg/ipv4 v0.0.0
)

replace github.com/srleyva/tcp-go/pkg/ipv4 => ./pkg/ipv4