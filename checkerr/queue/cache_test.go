package queue

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrMsgQItem_Push(t *testing.T) {
	s := New()
	s.Push(errors.New("hello"))
	s.Push(errors.New("hello2"))
	s.Push(errors.New("hello3"))
	s.Traverse()
	d := s.GetErrMQ()
	fmt.Println(d)
	s.Traverse()
}
