// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"testing"
	"container/vector"
)


func NanoSpec(gotest *testing.T, spec func(Context)) {
	c := newContext(gotest, spec)
	c.Run()
}

type Context interface {
	Specify(name string, body func())
}


type runContext struct {
	root         *aSpec
	current      *aSpec
	backtracking bool
}

func newContext(gotest *testing.T, spec func(Context)) *runContext {
	c := &runContext{}
	root := newSpec(nil, "<root>", func() { spec(c) })
	c.root = root
	c.current = root
	return c
}

func (this *runContext) Run() {
	safetyLimit := 10000 // just in case this program gets stuck in an infinite loop during development
	for this.root.ShouldExecute() && safetyLimit > 0 {
		this.backtracking = false
		this.root.Execute()
		safetyLimit--
	}
}

func (this *runContext) Specify(name string, body func()) {
	this.enterSpec(name, body)
	this.processSpec()
	this.exitSpec()
}

func (this *runContext) enterSpec(name string, body func()) {
	child := this.current.EnterChild(name, body)
	this.current = child
}

func (this *runContext) processSpec() {
	if !this.backtracking && this.current.ShouldExecute() {
		this.current.Execute()
		this.backtracking = true
	}
}

func (this *runContext) exitSpec() {
	this.current = this.current.Parent
}


type aSpec struct {
	Parent                *aSpec
	name                  string
	body                  func()
	children              *vector.Vector
	childrenSeenOnThisRun int
	hasBeenExecuted       bool
}

func newSpec(parent *aSpec, name string, body func()) *aSpec {
	return &aSpec{parent, name, body, new(vector.Vector), 0, false}
}

func (this *aSpec) ShouldExecute() bool {
	return !this.hasBeenExecuted
}

func (this *aSpec) Execute() {
	this.childrenSeenOnThisRun = 0
	this.body()
	this.hasBeenExecuted = this.allChildrenHaveBeenExecuted()
}

func (this *aSpec) allChildrenHaveBeenExecuted() bool {
	for _, v := range *this.children {
		child := v.(*aSpec)
		if !child.hasBeenExecuted {
			return false
		}
	}
	return true
}

func (this *aSpec) EnterChild(name string, body func()) *aSpec {
	this.childrenSeenOnThisRun++
	isUnseen := this.childrenSeenOnThisRun > this.children.Len()

	if isUnseen {
		child := newSpec(this, name, body)
		this.children.Push(child)
		return child
	}

	index := this.childrenSeenOnThisRun - 1
	child := this.children.At(index).(*aSpec)
	return child
}
