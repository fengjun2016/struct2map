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

func checkTwoSliceIsTheSame(selectAttributesA, selectAttributesB []interface{}) bool {
	isTheSame := true

	//check is the same length
	if len(selectAttributesA) != len(selectAttributesB) {
		return false
	}

	checkMap := make(map[string]interface{})
	for _, aSelect := range selectAttributesA {
		checkMap[aSelect] = true
	}

	//check is exists select a
	for _, bSelect := range selectAttributesB {
		if _, ok := checkMap[bSelect]; !ok {
			isTheSame = false
			break
		}
	}

	return isTheSame
}

func checkIsSameSelectAttributes(selectA, selectB []map[string]interface{}) bool {
	isTheSame := true

	//检查定制规格所选的规格类别数是否相等
	if len(selectA) != len(selectB) {
		return false
	}

	//分别获取两个规格选择类别的名称的[]string
	selectClassA := make([]string, 0)
	selectClassB := make([]string, 0)
	for _, aSelect := range selectA {
		selectClassA = append(selectClassA, aSelect["name"].(string))
	}

	for _, bSelect := range selectB {
		selectClassB = append(selectClassB, bSelect["name"].(string))
	}

	//检查两个选择规格类别是否名称一样
	if !CheckTwoSliceIsTheSame(selectClassA, selectClassB) {
		return false
	}

	//检查同一个规格类别选择的定制项个数和定制项是否一致
	for _, aSelect := range selectA {
		for _, bSelect := range selectB {
			if aSelect["name"] == bSelect["name"] {
				// check 同一个名称的规格选择项是否相同
				if !checkTwoSliceIsTheSame(aSelect["select_values"].([]string), bSelect["select_values"].([]string)) {
					return false
				}
			}
		}
	}
	return isTheSame
}
