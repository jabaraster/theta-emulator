package util

import (
    "reflect"
    "runtime"
)

func GetFunctionName(function interface{}) string {
    fv := reflect.ValueOf(function)
    return runtime.FuncForPC(fv.Pointer()).Name()
}
