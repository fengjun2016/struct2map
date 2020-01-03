package struct2map

import "reflect"

func obj2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if v.Field(i).CanInterface() && t.Field(i).Tag.Get("json") != "-" {
			data[t.Field(i).Tag.Get("json")] = v.Field(i).Interface()
		}
	}
	return data
}
