module github.com/srleyva/tcp-go

go 1.12

require (
	github.com/songgao/water v0.0.0-20190725173103-fd331bda3f4b
	github.com/srleyva/tcp-go/pkg/icmp v0.0.0
	github.com/srleyva/tcp-go/pkg/ipv4 v0.0.0
	golang.org/x/sys v0.0.0-20190804053845-51ab0e2deafa // indirect
)

replace github.com/srleyva/tcp-go/pkg/icmp => ./pkg/icmp

replace github.com/srleyva/tcp-go/pkg/ipv4 => ./pkg/ipv4
