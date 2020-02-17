package web

import (
	"sync"
	"syscall/js"
)

type Promise struct {
	Value
}

func (promise Promise) Catch(handler func(reason js.Value)) {
	wrapper := func(then js.Value, args []js.Value) interface{} {
		handler(args[0])
		return nil
	}
	promise.Call("catch", js.FuncOf(wrapper))
}

func (promise Promise) Then(handler func(value js.Value)) {
	wrapper := func(then js.Value, args []js.Value) interface{} {
		handler(args[0])
		return nil
	}
	promise.Call("then", js.FuncOf(wrapper))
}

func (promise Promise) Get() (msg Value, err Value) {
	// we'll wait only for one (the first) handler
	wg := sync.WaitGroup{}
	wg.Add(1)

	// register error handler
	catch := func(value js.Value) {
		err = Value{Value: value}
		wg.Done()
	}
	promise.Catch(catch)

	// register succsess handler
	then := func(reason js.Value) {
		msg = Value{Value: reason}
		wg.Done()
	}
	promise.Then(then)

	// wait until any handler is done
	wg.Wait()
	return msg, err
}
