package golib

import (
	"fmt"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

type Fun func()

const MAXGOCOUNT = 100

var Wg sync.WaitGroup //
var goCount int32     //当前协程数量
var ChanFun chan Fun

func init() {
	ChanFun = make(chan Fun, 10)
}

func Go(fun Fun) {
	var c int32
	select {
	case ChanFun <- fun:
		{
			return
		}
	default:
		{
			c = atomic.AddInt32(&goCount, 1)
			if c == MAXGOCOUNT {
				atomic.AddInt32(&goCount, -1)
			}
		}
	}

	Wg.Add(1)
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println(err)
				fmt.Println(debug.Stack())
			}
		}()
		fun()
		if c >= MAXGOCOUNT {
			return
		}
		for {
			select {
			case f := <-ChanFun:
				f()
			}
		}
		Wg.Done()
		atomic.AddInt32(&goCount, -1)
	}()
}
