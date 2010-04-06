// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"testing"
	"container/vector"
)


func TestAllSpecs(t *testing.T) {
	Run(t, StackSpec)
}

func StackSpec(c Context) {
	stack := new(vector.Vector)

	c.Specify("An empty stack", func() {

		c.Specify("contains no elements", func() {
			c.Expect(stack.Len()).Equals(0)
		})
	})

	c.Specify("When elements are pushed to a stack", func() {
		stack.Push("first push")
		stack.Push("last push")

		c.Specify("then it contains some elements", func() {
			c.Expect(stack.Len()).NotEquals(0)
		})
		c.Specify("the element pushed last is popped first", func() {
			x := stack.Pop()
			c.Expect(x).Equals("last push")
		})
		c.Specify("the element pushed first is popped last", func() {
			stack.Pop()
			x := stack.Pop()
			c.Expect(x).Equals("first push")
		})
	})
}
