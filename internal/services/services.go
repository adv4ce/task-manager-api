package services

import (
	"reflect"
	"fmt"
	"strings"
)

func SturcToMap(s *PatchTask, chgVal *map[string]string) {
	val := reflect.ValueOf(s).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		if value.IsZero() {
			continue
		}
		
		jsonTag := field.Tag.Get("json")
        if jsonTag != "" && jsonTag != "id"{
            fieldName := strings.Split(jsonTag, ",")[0]
            (*chgVal)[fieldName] = fmt.Sprintf("%v", value.Interface())
        }
	}
}