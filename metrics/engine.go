package metrics

import (
	"runtime"
	"time"
)

type Metrics struct {
	FPS			int
	FrameTime	time.Duration
	Memory		uint64
	DrawCalls	int
	Objects		int
}

type Monitor struct {
	metricsChan chan Metrics
	lastUpdate time.Time
	frameCount int
}

func NewMonitor() *Monintor {
	m := &Monitor {
		metricsChan: make(chan Metrics, 1)
	}
	go m.run()
	return m
}

func (m *Monintor) run(){
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		m.metricsChan <- Metrics {
			FPS:	m.frameCount,
			FrameTime:	time.Since(m.lastUpdate) / time.Duration(m.frameCount),
			Memory:		memStats.Alloc,
		}
		m.frameCount = 0
	}
}

func (m *Monitor) RecordFrame() {
	m.frameCount++
	m.lastUpdate = time.Now()
}