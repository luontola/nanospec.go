// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"testing"
)


func TT(t *testing.T) *TestTool {
	return &TestTool{t}
}

type TestTool struct {
	t *testing.T
}

func (this *TestTool) AssertEquals(expected, actual interface{}) {
	if expected != actual {
		this.t.Errorf("Expected '%v', got '%v'", expected, actual)
	}
}
