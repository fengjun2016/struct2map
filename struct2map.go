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

func addAttr(obj interface{}, k string, v interface{}) map[string]interface{} {
	mapData := obj2Map(obj)
	mapData[k] = v
	return mapData
}

func deleteAttr(obj interface{}, k string) map[string]interface{} {
	mapData := obj2Map(obj)
	delete(mapData, k)
	return mapData
}

func updateAttr(obj interface{}, k string, v interface{}) map[string]interface{} {
	mapData := obj2Map(obj)
	mapData[k] = v
	return mapData
}

func mergeObj(objA, objB interface{}, mapRes *map[string]interface{}) {
	mapRes = obj2Map(objA)
	mapRes = obj2Map(objB)
	return
}
