package utils

import "reflect"

func IsEmptyValue(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.String:
		return val.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint() == 0
	case reflect.Slice:
		return val.Len() == 0
	}
	return false
}
