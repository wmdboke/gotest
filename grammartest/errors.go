package grammartest

import (
	"errors"
	"log"

	pkgerr "github.com/pkg/errors"
)

func Errors_test() {
	// 通过 deep 查看调用堆栈
	log.Printf("%+v", deep3()) //with stack
	err2 := msg2()
	log.Printf("%+v", err2)     //no stack
	log.Printf("=============") //分隔符
	log.Printf("%+v", pkgerr.WithStack(err2))
}
func deep3() error {
	return deep2()
}

func deep2() error {
	return pkgerr.Wrap(deep1(), "deep2")
}

func deep1() error {
	return errors.New("deep1")
}

func msg() error {
	return errors.New("msg")
}

func msg2() error {
	return pkgerr.WithMessage(msg(), "msg2")

}
