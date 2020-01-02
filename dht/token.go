package dht

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"math/big"
	"math/rand"
	"time"
)

type Id []byte
type Nid *big.Int

func GenerateId() Id  {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	h := sha1.New()
	io.WriteString(h,time.Now().String())
	io.WriteString(h,string(random.Int()))
	return h.Sum(nil)
}

//Neighbor get neighbor
func (id Id) Neighbor(tableID Id) Id {
	return append(id[:6], tableID[6:]...)
}

//Int get int
func (id Id) Int() *big.Int {
	return big.NewInt(0).SetBytes(id)
}

func (id Id) String() string {
	return hex.EncodeToString(id)
}
