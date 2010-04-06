// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
)


const runtimeCallerBugfix = 1

func callerLocation() string {
	if _, file, line, ok := runtime.Caller(2 + runtimeCallerBugfix); ok {
		return fmt.Sprintf("%v:%v", filename(file), line)
	}
	return "<unknown file>"
}

func filename(fullpath string) string {
	_, file := path.Split(fullpath)
	return file
}

func functionName(function interface{}) string {
	fval := reflect.NewValue(function).(*reflect.FuncValue)
	if f := runtime.FuncForPC(fval.Get()); f != nil {
		return f.Name()
	}
	return "<unknown function>"
}
