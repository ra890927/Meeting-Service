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

func OverwriteValue(base, edit interface{}) interface{} {
	va := reflect.ValueOf(base)
	vb := reflect.ValueOf(edit)

	if va.Kind() != reflect.Ptr || vb.Kind() != reflect.Ptr {
		panic("Both arguments must be pointers to struct")
	}

	va = va.Elem()
	vb = vb.Elem()

	res := reflect.New(va.Type()).Elem()
	res.Set(va)

	for i := 0; i < va.NumField(); i++ {
		fieldA := res.Field(i)
		fieldB := vb.Field(i)

		if !IsEmptyValue(fieldB) {
			fieldA.Set(fieldB)
		}
	}

	return res.Addr().Interface()
}
