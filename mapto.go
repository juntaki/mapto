package mapto

import (
	"reflect"
	"strings"
)

const key = "mapto"

func Map(dest interface{}, src interface{}) {
	if reflect.ValueOf(dest).Kind() != reflect.Ptr ||
		reflect.ValueOf(src).Kind() != reflect.Ptr ||
		reflect.Indirect(reflect.ValueOf(dest)).Kind() != reflect.Struct ||
		reflect.Indirect(reflect.ValueOf(src)).Kind() != reflect.Struct {
		panic("invalid input type input must be pointer of struct")
	}

	v := reflect.Indirect(reflect.ValueOf(dest))
	for i := 0; i < v.NumField(); i++ {
		tagArray := strings.Split(v.Type().Field(i).Tag.Get(key), ".")
		srcVal := reflect.ValueOf(src)
		destVal := reflect.ValueOf(dest).Elem().Field(i)

		for _, tag := range tagArray {
			switch srcVal.Kind() {
			case reflect.Ptr:
				srcVal = srcVal.Elem().FieldByName(tag)
			case reflect.Struct:
				srcVal = srcVal.FieldByName(tag)
			case reflect.Invalid:
				break
			default:
				panic(srcVal.Kind())
			}
		}
		if !srcVal.IsValid() {
			continue
		}

		if srcVal.Kind() != destVal.Kind() {
			panic("invalid mapping")
		}

		copy(destVal, srcVal)
	}
}

func copy(dest reflect.Value, src reflect.Value) {
	switch src.Kind() {
	case reflect.Bool:
		dest.SetBool(src.Bool())
	case reflect.Int:
		dest.SetInt(src.Int())
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.Uintptr:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Complex64:
	case reflect.Complex128:
	case reflect.Map:
	case reflect.Ptr:
		if !dest.Elem().IsValid() {
			dest.Set(reflect.New(src.Elem().Type()))
		}
		copy(dest.Elem(), src.Elem())
	case reflect.Slice:
		fallthrough
	case reflect.Array:
		reflect.Copy(dest, src)
	case reflect.String:
		dest.SetString(src.String())
	case reflect.Struct:
		fallthrough
	case reflect.Chan:
		fallthrough
	case reflect.Func:
		fallthrough
	case reflect.Interface:
		fallthrough
	case reflect.UnsafePointer:
		panic("don't use for domain object")
	}
}
