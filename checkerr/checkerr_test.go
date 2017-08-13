package checkerr

import (
	"testing"

	"fmt"

	"github.com/servicecomb/service-center/util/errors"
)

func TestCheckErr(t *testing.T) {
	o := Check(func() (out string, err error) {
		out = "hello world"
		err = errors.New("hello error")
		return
	}())
	fmt.Println(o)
}
