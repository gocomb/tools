package queue
//消息定义
type ErrMsgQ struct {
	msg     error
	key     int
	nextPtr *ErrMsgQ
	prevPtr *ErrMsgQ
}
//消息队列定义
type ErrMsgQItem struct {
	Top    *ErrMsgQ
	Bottom *ErrMsgQ
}

