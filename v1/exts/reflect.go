package exts

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var (
	ErrStructFieldNotExist = errors.New("struct field not exist")
)

func Reflect_ParseString(value string, kind reflect.Kind) (reflect.Value, error) {
	var v any
	var err error

	switch kind {
	case reflect.Bool:
		v, err = strconv.ParseBool(value)
	case reflect.String:
		v = value
	case reflect.Int:
		v, err = strconv.Atoi(value)
	default:
		return reflect.Value{}, fmt.Errorf("not support kind %v", kind)
	}

	if err != nil {
		return reflect.Value{}, err
	}

	return reflect.ValueOf(v), nil
}

func Reflect_New(t reflect.Type) reflect.Value {
	switch t.Kind() {
	case reflect.Slice:
		return reflect.MakeSlice(t, 0, 0)
	case reflect.Struct:
		return reflect.New(t)
	default:
		return reflect.New(t)
	}
}

func Reflect_IsNil(v reflect.Value) bool {
	switch v.Type().Kind() {
	case reflect.Chan:
	case reflect.Func:
	case reflect.Interface:
	case reflect.Map:
	case reflect.Pointer:
	case reflect.Slice:
		return v.IsNil()
	}

	return false
}

func Reflect_IsArray(t reflect.Type) bool {
	typeKind := t.Kind()

	if typeKind == reflect.Pointer {
		typeKind = t.Elem().Kind()
	}

	if typeKind == reflect.Array || typeKind == reflect.Slice {
		return true
	}

	return false
}
