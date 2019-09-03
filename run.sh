#!/bin/sh

go build main.go
cp main /tmp/main
setcap cap_net_bind_service=+ep /tmp/main
/tmp/main &
pid=$!

echo $pid

ip addr add 192.168.0.1/24 dev tun0
ip link set up dev tun0

trap "kill $pid" INT TERM
wait $pid
