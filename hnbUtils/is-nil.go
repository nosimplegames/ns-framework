package hnbUtils

import "reflect"

func IsNil(x Any) bool {
	return x == nil || (reflect.ValueOf(x).Kind() == reflect.Ptr && reflect.ValueOf(x).IsNil())
}
