# Error handling agent tools
## overview
* This module design is for golang error handling. It service the agency as the output of the function return value and convert to
inteface slice, when found the error types, triggering error handling
* 本模块设计是为了方便golang的错误处理，将函数的中间输出转到一个inteface切片的中间层，代理轮询中间层的
类型，当发现error类型时，触发错误处理机制
## todo
加入网络通知模块
## 使用
安装
``````
go get -u github.com/gocomb/tools
import "github.com/gocomb/tools/checkerr"

``````
代理函数输出
````
o := checkerr.Check(func() (out string, err error) {
		out = "hello world"
		err = errors.New("hello error")
		return
	}())
````
打印错误
````
out:=o.Do("print")
````
查看函数返回值
````
fmt.Println(out[0])
````
painc错误
````
o.Do("painc")
````
压入错误消息队列
````
o.Do("push")
````
从队列中取出一个错误
````
checkerr.GetErr()
````
