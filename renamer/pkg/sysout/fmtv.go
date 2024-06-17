package sysout

import (
	"fmt"
	"reflect"
)

func _isstringer(v interface{}) bool {
	_, ok := v.(fmt.Stringer)
	return ok
}

func _isslice(v interface{}) bool {
	isslice := reflect.TypeOf(v).Kind() == reflect.Slice
	isarray := reflect.TypeOf(v).Kind() == reflect.Array
	return isslice || isarray
}

func _isstruct(f interface{}) bool {
	switch v := reflect.ValueOf(f); v.Kind() {
	case reflect.Ptr:
		_isstruct(v.Elem()) // dereference the pointer
	case reflect.Struct:
		return true
	}
	return false
}

func _fmtslice(v interface{}) string {
	strs := []interface{}{}
	count := reflect.ValueOf(v).Len()
	for i := 0; i < count; i++ {
		element := reflect.ValueOf(v).Index(i).Interface()
		strs = append(strs, fmtv(element))
	}
	return tojson(strs)
}

func fmtv(v interface{}) interface{} {
	switch {
	case v == nil:
		return ""
	case _isstringer(v):
		return v.(fmt.Stringer).String()
	case _isslice(v):
		return _fmtslice(v)
	case _isstruct(v):
		return tojson(v)
	default:
		return v
	}
}
