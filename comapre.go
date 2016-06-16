package utils

import (
	"fmt"
	"sort"

    "github.com/garyburd/redigo/redis"
)

func CompareInterface(value1, value2 interface{}) bool {
	switch value1Type := value1.(type) {
		case []interface{}:
			value1Array := make([]string, 0)
			value2Array := make([]string, 0)
			value1Values, _ := redis.Values(value1, nil)
			value2Values, err := redis.Values(value2, nil)
			if err != nil {
				return false
			}
			if len(value1Values) != len(value2Values) {
				return false
			}
			for ind, item := range value1Values {
				strItem1, err := redis.String(item, nil)
				if err != nil {
					return false
				}
				strItem2, err := redis.String(value2Values[ind], nil)
				if err != nil {
					return false
				}
				value1Array = append(value1Array, strItem1)
				value2Array = append(value2Array, strItem2)
			}
			sort.Strings(value1Array)
			sort.Strings(value2Array)
			for ind, strValue := range value1Array {
				if strValue != value2Array[ind] {
					return false
				}
			}
		case []byte:
			strValue1, _   := redis.String(value1, nil)
			strValue2, err := redis.String(value2, nil)
			if err != nil {
				return false
			}
			if strValue1 != strValue2 {
				return false
			}
        case string:
            strValue1 := value1.(string)
            strValue2 := value2.(string)
            if strValue1 != strValue2 {
                return false
            }
        case []string:
            return CompareStringSlice(value1.([]string), value2.([]string))
		default:
			fmt.Println("Unexpect value type: ", value1Type)
			return false
	}
	return true
}

func CompareStringSlice(value1, value2 []string) bool {
	if len(value1) != len(value2) {
		return false
	}
    sort.Strings(value1)
    sort.Strings(value2)
	for i, val := range value1 {
		if val != value2[i] {
			return false
		}
	}
	return true
}
