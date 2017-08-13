package checkerr

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/gocomb/tools/checkerr/queue"
)

var msgQ = queue.New()

//入口函数
func Check(args ...interface{}) CheckErr {
	errCheck := CheckErr{}
	errCheck.val = args
	return errCheck
}

//根据method执行错误处理方式，以切片形式返回原函数返回值
func (p *CheckErr) Do(method string) []interface{} {
	val := p.val.([]interface{})
	switch method {
	case "print":
		printErr(val)
	case "painc":
		panicErr(val)
	case "push":
		push(val)
	}
	return val
}

//打印错误
func printErr(args []interface{}) {
	for _, v := range args {
		if v != nil {
			if strings.Contains(reflect.TypeOf(v).String(), "error") {
				fmt.Printf("%c[1;31m%s[ERROR]\t%s%c[0m\n\n", 0x1B, time.Now().Format("2006/01/02 - 15:04:05"), v, 0x1B)
			}
		}
	}
}

//painc错误
func panicErr(args []interface{}) {
	for _, v := range args {
		if v != nil {
			if strings.Contains(reflect.TypeOf(v).String(), "error") {
				panic(v)
			}
		}
	}
}

//加入error队列
func push(args []interface{}) {
	for _, v := range args {
		if v != nil {
			if strings.Contains(reflect.TypeOf(v).String(), "error") {
				err := v.(error)
				msgQ.Push(err)
			}
		}
	}

}

//从队列中取出error
func GetErr() error {
	g := msgQ.GetErrMQ()
	return g.FetchErr()
}
