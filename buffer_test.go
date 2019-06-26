package ds

import (
	"testing"
	"time"
)

func TestBuffer(t *testing.T) {
	b := NewBuffer(5, 1, time.Millisecond*100)
	go func() {
		for {
			select {
			case _, ok := <-b.C:
				if !ok {
					return
				}
			}
		}
	}()

	for i := 0; i < 1024; i++ {
		b.Push(i)
	}

	for i := 0; i < 4; i++ {
		b.Push(i)
	}

	time.Sleep(time.Second)

	for i := 0; i < 3; i++ {
		b.Push(i)
	}

	b.Close()
}
