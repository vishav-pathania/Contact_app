package utils

import "reflect"

func GetVariableType(v interface{}) string {
	return reflect.TypeOf(v).String()
}
