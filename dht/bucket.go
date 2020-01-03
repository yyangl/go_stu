package dht

import (
	"fmt"
	"math/big"
	"sync"
	"time"
)

const (
	maxNode = 8
)

type Bucket struct {
	sync.RWMutex
	// 最小node ID数字值
	min *big.Int
	// 最大node ID数字值
	max *big.Int
	// Node 节点
	nodes []*Node
	// 最后活跃时间
	lastActivity time.Time
}

func NewBucket(min, max *big.Int) *Bucket {
	fmt.Printf("min = %v,max = %v\n", min, max)
	return &Bucket{
		RWMutex:      sync.RWMutex{},
		min:          min,
		max:          max,
		nodes:        make([]*Node, 8),
		lastActivity: time.Now(),
	}
}

func (b *Bucket) Length() int {
	return len(b.nodes)
}
