package collection

import (
	"gin-essential/pkg/timex"
	"sync"
	"time"
)

type (
	// RollingWindowOption ..
	RollingWindowOption func(*RollingWindow)
	// RollingWindow ..
	RollingWindow struct {
		lock          sync.RWMutex
		size          int
		win           *window
		interval      time.Duration
		offset        int
		ignoreCurrent bool
		lastTime      time.Duration
	}
)

// NewRollingWindow returns a RollingWindow that with size buckets and time interval,
// use opts to customize the RollingWindow.
func NewRollingWindow(size int, interval time.Duration, opts ...RollingWindowOption) *RollingWindow {
	if size < 1 {
		panic("size must be greater than 0")
	}

	w := &RollingWindow{
		size:     size,
		win:      newWindow(size),
		interval: interval,
		lastTime: timex.Now(),
	}

	for _, opt := range opts {
		opt(w)
	}
	return w
}

// Reduce runs fn on all buckets, ignore current bucket if ignoreCurrent was set.
func (rw *RollingWindow) Reduce(fn func(b *Bucket)) {
	rw.lock.RLock()
	defer rw.lock.RUnlock()

	var diff int
	span := rw.span()
	// ignore current bucket, because of partial data
	if span == 0 && rw.ignoreCurrent {
		diff = rw.size - 1
	} else {
		diff = rw.size - span
	}
	if diff > 0 {
		offset := (rw.offset + span + 1) % rw.size
		rw.win.reduce(offset, diff, fn)
	}
}

// Add adds value to current bucket.
func (rw *RollingWindow) Add(v float64) {
	rw.lock.Lock()
	defer rw.lock.Unlock()
	rw.updateOffset()
	rw.win.add(rw.offset, v)
}

func (rw *RollingWindow) updateOffset() {
	span := rw.span()
	if span <= 0 {
		return
	}

	offset := rw.offset
	for i := 0; i < span; i++ {
		rw.win.resetBucket((offset + 1 + i) % rw.size)
	}

	rw.offset = (offset + span) % rw.size
	now := timex.Now()
	// ????
	rw.lastTime = now - (now-rw.lastTime)%rw.interval
}

func (rw *RollingWindow) span() int {
	offset := int(timex.Since(rw.lastTime) / rw.interval)
	if 0 <= offset && offset < rw.size {
		return offset
	}
	return rw.size
}

type window struct {
	buckets []*Bucket
	size    int
}

func newWindow(size int) *window {
	buckets := make([]*Bucket, size)
	for i := 0; i < size; i++ {
		buckets[i] = new(Bucket)
	}

	return &window{
		buckets: buckets,
		size:    size,
	}
}

func (w *window) resetBucket(offset int) {
	w.buckets[offset%w.size].reset()
}

func (w *window) add(offset int, v float64) {
	w.buckets[offset%w.size].add(v)
}

func (w *window) reduce(start, count int, fn func(b *Bucket)) {
	for i := 0; i < count; i++ {
		fn(w.buckets[(start+i)%w.size])
	}
}

// Bucket defines the bucket that holds sum and num of additions.
type Bucket struct {
	Sum   float64
	Count int64
}

func (b *Bucket) reset() {
	b.Sum = 0
	b.Count = 0
}

func (b *Bucket) add(v float64) {
	b.Sum += v
	b.Count++
}
