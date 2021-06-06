package pragram

import (
	"context"
	"fmt"
)

type Selectx struct {
	Ctx context.Context
	Msg chan int
}

func (s *Selectx) SelectTest() {
	defer fmt.Println("exit")
	for {
		select {
		case m := <-s.Msg:
			fmt.Println(m)
		case <-s.Ctx.Done():
			fmt.Println("s.ctx.done")
			return
		}
	}
}
