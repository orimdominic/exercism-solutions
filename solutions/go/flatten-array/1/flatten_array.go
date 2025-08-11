package flatten

import "reflect"

var list []interface{} = make([]interface{}, 0)

func Flatten(nested interface{}) []interface{} {

	t := reflect.TypeOf(nested)
	if t.Kind() == reflect.Slice {
		for _, v := range nested.([]interface{}) {
			if v == nil {
				continue
			}

			if reflect.TypeOf(v).Kind() == reflect.Slice {
				innerList := Flatten(v)
				list = append(list, innerList...)
				continue
			}

			list = append(list, v)
		}
	}

	ret := list
	list = make([]interface{}, 0)
	return ret
}
