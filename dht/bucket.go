package dht

import (
	"sync"
	"time"
)

const (
	maxNode = 8
)

type Bucket struct {
	sync.RWMutex
	// 最小node ID数字值
	min int
	// 最大node ID数字值
	max int
	// Node 节点
	nodes []*Node
	// 最后活跃时间
	lastActivity time.Time
}


