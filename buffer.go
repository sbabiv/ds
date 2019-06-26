package ds

import (
	"sync"
	"time"
)

type Buffer struct {
	sync.Mutex
	t     *time.Timer
	data  []interface{}
	delay time.Duration
	C     chan []interface{}
}

func NewBuffer(chunkSize, bufferSize int, delay time.Duration) *Buffer {
	buffer := &Buffer{
		data:  make([]interface{}, 0, chunkSize),
		C:     make(chan []interface{}, bufferSize),
		t:     time.NewTimer(delay),
		delay: delay,
	}

	go buffer.collect()

	return buffer
}

func (b *Buffer) Push(item interface{}) {
	b.Lock()
	defer b.Unlock()

	b.data = append(b.data, item)
	if len(b.data) == cap(b.data) {
		b.send()
	}
}

func (b *Buffer) send() {
	b.t.Stop()
	defer b.t.Reset(b.delay)

	tmp := make([]interface{}, len(b.data))
	copy(tmp, b.data)
	b.C <- tmp
	b.data = b.data[:0]
}

func (b *Buffer) collect() {
	b.Lock()
	defer b.Unlock()

	if len(b.data) == 0 {
		return
	}
	b.send()
}

func (b *Buffer) Close() {
	b.t.Stop()
	b.collect()
	close(b.C)
}
