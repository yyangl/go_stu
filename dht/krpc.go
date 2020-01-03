package dht

import "github.com/zeebo/bencode"

type Krpc struct {
	t string
	y string
	q string
	a interface{}
}

func (k *Krpc) Encode() (string, error) {
	return bencode.EncodeString(k)
}

func (k *Krpc) SetArgs(args interface{}) {
	k.a = args
}
