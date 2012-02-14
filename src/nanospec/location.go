// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"fmt"
	"reflect"
	"runtime"
)

func callerLocation() string {
	if _, file, line, ok := runtime.Caller(2); ok {
		return fmt.Sprintf("%v:%v", file, line)
	}
	return "<unknown file>"
}

func functionName(function interface{}) string {
	fval := reflect.ValueOf(function)
	if f := runtime.FuncForPC(fval.Pointer()); f != nil {
		return f.Name()
	}
	return "<unknown function>"
}
