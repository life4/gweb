package main

import (
	"sync"
)

type State struct {
	Stop SubState
}

type SubState struct {
	Requested bool
	done      bool
	wg        sync.WaitGroup
}

func (state *SubState) Request() {
	state.Requested = true
	state.done = false
	state.wg = sync.WaitGroup{}
	state.wg.Add(1)
}

func (state *SubState) Done() {
	state.done = true
	state.wg.Done()
}

func (state *SubState) Wait() {
	state.wg.Wait()
}
