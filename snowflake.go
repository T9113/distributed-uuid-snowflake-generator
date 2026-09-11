package snowflake

import (
	"sync"
	"time"
)

type Generator struct {
	mu        sync.Mutex
	lastTime  int64
	workerID  int64
	sequence  int64
}

func NewGenerator(workerID int64) *Generator {
	return &Generator{workerID: workerID}
}

func (g *Generator) NextID() int64 {
	g.mu.Lock()
	defer g.mu.Unlock()
	now := time.Now().UnixMilli()
	if now == g.lastTime {
		g.sequence = (g.sequence + 1) & 4095
	} else {
		g.sequence = 0
		g.lastTime = now
	}
	return (now << 22) | (g.workerID << 12) | g.sequence
}
