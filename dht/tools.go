package dht

import "net"

func ParseByteStream(data []byte) []*Node {
	var nodes []*Node
	for j := 0 ;j < len(data); j = j+26 {
		if j + 26 > len(data) {
			break
		}
		kn := data[j:j+26]
		node := new(Node)
		node.id = kn[:20]
		addr := new(net.UDPAddr)
		addr.IP = kn[20:24]
		port := kn[24:26]
		addr.Port = int(port[0])<<8 + int(port[1])
		nodes = append(nodes, node)
	}
	return nodes
}
