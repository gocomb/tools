package queue

import (
	"fmt"
	"unsafe"
)

//新建错误队列
func New() (m *ErrMsgQItem) {
	m = new(ErrMsgQItem)
	*m = ErrMsgQItem{Top: new(ErrMsgQ), Bottom: new(ErrMsgQ)}
	*(m.Top) = ErrMsgQ{key: 0, nextPtr: m.Bottom, prevPtr: nil}
	*(m.Bottom) = ErrMsgQ{key: 0, nextPtr: nil, prevPtr: m.Top}
	return
}

//压入新消息
func (m *ErrMsgQItem) Push(data error) {
	queueTemp := new(ErrMsgQ)
	*queueTemp = ErrMsgQ{
		key:     m.Bottom.key + 1,
		msg:     data,
		nextPtr: nil,
		prevPtr: new(ErrMsgQ),
	}
	temp := m.Bottom
	temp.nextPtr = queueTemp
	queueTemp.prevPtr = temp
	m.Bottom = queueTemp
}

//遍历消息
func (m *ErrMsgQItem) Traverse() {
	temp := m.Top
	fmt.Println(temp)
	condition := true
	for condition {
		temp = temp.nextPtr
		fmt.Println(unsafe.Pointer(temp), temp)
		if temp == nil {
			condition = false
		}
	}
}

//取出并删除消息
func (m *ErrMsgQItem) GetErrMQ() ErrMsgQ {
	out := *m.Bottom
	temp := m.Bottom.prevPtr
	m.Bottom.prevPtr.nextPtr = nil
	m.Bottom = temp
	return out
}

//得到error
func (err ErrMsgQ) FetchErr() error {
	return err.msg
}
