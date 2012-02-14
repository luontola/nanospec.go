// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"container/list"
	"testing"
)

func TestAllSpecs(t *testing.T) {
	Run(t, StackSpec)
}

func StackSpec(c Context) {
	stack := new(list.List)

	c.Specify("An empty stack", func() {

		c.Specify("contains no elements", func() {
			c.Expect(stack.Len()).Equals(0)
		})
	})

	c.Specify("When elements are pushed onto a stack", func() {
		stack.PushFront("pushed first")
		stack.PushFront("pushed last")

		c.Specify("then it contains some elements", func() {
			c.Expect(stack.Len()).NotEquals(0)
		})
		c.Specify("the element pushed last is popped first", func() {
			poppedFirst := stack.Remove(stack.Front())
			c.Expect(poppedFirst).Equals("pushed last")
		})
		c.Specify("the element pushed first is popped last", func() {
			stack.Remove(stack.Front())
			poppedLast := stack.Remove(stack.Front())
			c.Expect(poppedLast).Equals("pushed first")
		})
	})
}
