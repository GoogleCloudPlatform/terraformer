package ali_mns

import (
	"sync/atomic"
	"time"
)

type QPSMonitor struct {
	qpsLimit     int32
	latestIndex  int32
	delaySecond  int32
	totalQueries []int32
}

func (p *QPSMonitor) Pulse() {
	index := int32(time.Now().Second()) % p.delaySecond

	if p.latestIndex != index {
		atomic.StoreInt32(&p.latestIndex, index)
		atomic.StoreInt32(&p.totalQueries[p.latestIndex], 0)
	}

	atomic.AddInt32(&p.totalQueries[index], 1)
}

func (p *QPSMonitor) QPS() int32 {
	var totalCount int32 = 0
	for i, queryCount := range p.totalQueries {
		if int32(i) != p.latestIndex {
			totalCount += queryCount
		}
	}
	return totalCount / (p.delaySecond - 1)
}

func (p *QPSMonitor) checkQPS() {
	p.Pulse()
	if p.qpsLimit > 0 {
		for p.QPS() > p.qpsLimit {
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func NewQPSMonitor(delaySecond int32, qpsLimit int32) *QPSMonitor {
	if delaySecond < 5 {
		delaySecond = 5
	}
	monitor := QPSMonitor{
		qpsLimit:     qpsLimit,
		delaySecond:  delaySecond,
		totalQueries: make([]int32, delaySecond),
	}
	return &monitor
}
