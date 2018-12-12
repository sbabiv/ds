package ds

import (
	"testing"
)

func TestList(t *testing.T) {
	list := NewList()

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	if list.Len() != 10 {
		t.Fatal("list len != 10")
	}

	list.Clear()
	if list.len != 0 {
		t.Fatal("list clear fail")
	}

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	for i := 0; i < 10; i++ {
		if ok, _ := list.Get(); !ok {
			t.Fatal("list get fail")
		}
	}

	if ok, value := list.Get(); ok {
		t.Fatalf("list get error %v\n", value.(int))
	}
}
