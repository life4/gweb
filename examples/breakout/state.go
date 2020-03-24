package main

import (
	"sync"
)

type State struct {
	Stop SubState
}

type SubState struct {
	Requested bool
	Completed bool
	wg        sync.WaitGroup
}

func (state *SubState) Request() {
	state.Requested = true
	state.Completed = false
	state.wg = sync.WaitGroup{}
	state.wg.Add(1)
}

func (state *SubState) Complete() {
	state.Completed = true
	state.wg.Done()
}

func (state *SubState) Wait() {
	state.wg.Wait()
}
