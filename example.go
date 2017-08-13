package main

import (
	"errors"
	"fmt"

	"github.com/gocomb/tools/checkerr"
)

func main() {
	testCheckErr()

}
func defPrint(args ...interface{}) {
	fmt.Println(args)
}
func noErr() error {
	return nil
}
func testCheckErr() {
	fmt.Printf("%c[1;32m%s%c[0m\n", 0x1B, "start checkErr tools", 0x1B)
	o := checkerr.Check(func() (out string, err error) {
		out = "hello world"
		err = errors.New("hello error")
		return
	}())
	out := o.Do("print")
	fmt.Println(out[1])
	ex := checkerr.Check(func() (out string, err error) {
		out = "hello world msq"
		err = errors.New("error:hello error msq")
		return
	}())
	ex.Do("push")
	fmt.Println(checkerr.GetErr())

	o.Def(defPrint)("test checkerr.def")
	o2 := checkerr.Check(noErr())
	o2.Def(defPrint)("test2 checkerr.def")
	fmt.Printf("%c[1;32m%s%c[0m\n", 0x1B, "checkErr os done, test painc error", 0x1B)
	o.Do("painc")
}
