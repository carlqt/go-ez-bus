package helpers

import "reflect"

func MapToStruct(input map[string]interface{}, output interface{}) {
	t := reflect.ValueOf(output).Elem()
	for k, v := range input {
		val := t.FieldByName(k)
		val.Set(reflect.ValueOf(v))
	}
}
