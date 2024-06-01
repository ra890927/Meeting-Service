package utils

import (
	"reflect"
	"time"
)

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
	// time.Time is a struct, so we need to check it separately
	case reflect.Struct:
		if val.Type().String() == "time.Time" {
			// zero value of time.Time is
			// "0001-01-01 00:00:00 +0000 UTC" (in time package)
			// time.Unix(0, 0) (in unix timestamp)
			return val.Interface().(time.Time).IsZero() || val.Interface().(time.Time).Unix() == 0
		}
	}
	return false
}

func OverwriteValue(base, edit interface{}) {
	vbase := reflect.ValueOf(base)
	vedit := reflect.ValueOf(edit)

	if vbase.Kind() != reflect.Ptr || vedit.Kind() != reflect.Ptr {
		panic("Both arguments must be pointers to struct")
	}

	vbase = vbase.Elem()
	vedit = vedit.Elem()

	for i := 0; i < vbase.NumField(); i++ {
		fieldBase := vbase.Field(i)
		fieldEdit := vedit.Field(i)

		if IsEmptyValue(fieldBase) {
			fieldBase.Set(fieldEdit)
		}
	}
}
