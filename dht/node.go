package dht

import (
	"net"
	"time"
)

type Node struct {
	//id *bitmap
	// udp地址
	addr *net.UDPAddr
	id Id
	// 上次活动的时间
	lastActiveTime time.Time
}

