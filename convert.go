package utils

import (
	"fmt"
    "errors"
	"strconv"
)

var (
    ErrNil = errors.New("convert: nil given")
    errNegativeInt = errors.New("convert: negative value given for Uint64")
)

func Int(data interface{}) (int, error) {
	switch data := data.(type) {
    case int:
        return data, nil
	case int64:
		x := int(data)
		if int64(x) != data {
			return 0, strconv.ErrRange
		}
		return x, nil
	case []byte:
		return strconv.Atoi(string(data))
    case string:
        return strconv.Atoi(string(data))
	case nil:
		return 0, ErrNil
	}
	return 0, fmt.Errorf("convert: unexpected type for Int, got type %T", data)
}

func Int64(data interface{}) (int64, error) {
	switch data := data.(type) {
    case int:
        return int64(data), nil
	case int64:
		return data, nil
	case []byte:
		return strconv.ParseInt(string(data), 10, 64)
    case string:
        return strconv.ParseInt(data, 10, 64)
	case nil:
		return 0, ErrNil
	}
	return 0, fmt.Errorf("convert: unexpected type for Int64, got type %T", data)
}

func Uint64(data interface{}) (uint64, error) {
	switch data := data.(type) {
    case int:
        if data < 0 {
            return 0, errNegativeInt
        }
	case int64:
		if data < 0 {
			return 0, errNegativeInt
		}
		return uint64(data), nil
    case uint64:
        return data, nil
	case []byte:
		return strconv.ParseUint(string(data), 10, 64)
    case string:
        return strconv.ParseUint(data, 10, 64)
	case nil:
		return 0, ErrNil
    }
	return 0, fmt.Errorf("convert: unexpected type for Uint64, got type %T", data)
}

func Float64(data interface{}) (float64, error) {
	switch data := data.(type) {
    case int:
        return float64(data), nil
    case int64:
        return float64(data), nil
    case float64:
        return data, nil
	case []byte:
		return strconv.ParseFloat(string(data), 64)
    case string:
        return strconv.ParseFloat(data, 64)
	case nil:
		return 0, ErrNil
	}
	return 0, fmt.Errorf("convert: unexpected type for Float64, got type %T", data)
}

func String(data interface{}) (string, error) {
	switch data := data.(type) {
    case bool:
        return strconv.FormatBool(data), nil
    case int:
        return strconv.Itoa(data), nil
    case int64:
        return strconv.FormatInt(data, 10), nil
    case uint64:
        return strconv.FormatUint(data, 10), nil
    case float32:
        return strconv.FormatFloat(float64(data), 'f', -1, 32), nil
    case float64:
        return strconv.FormatFloat(data, 'f', -1, 64), nil
	case []byte:
		return string(data), nil
	case string:
		return data, nil
	case nil:
		return "", ErrNil
	}
	return "", fmt.Errorf("convert: unexpected type for String, got type %T", data)
}

func Bytes(data interface{}) ([]byte, error) {
    retStr, err := String(data)
    return []byte(retStr), err  //TODO
}

func Bool(data interface{}) (bool, error) {
	switch data := data.(type) {
    case bool:
        return data, nil
    case int:
        return data != 0, nil
	case int64:
		return data != 0, nil
    case uint64:
        return data != 0, nil
	case []byte:
		return strconv.ParseBool(string(data))
    case string:
        return strconv.ParseBool(data)
	case nil:
		return false, ErrNil
	}
	return false, fmt.Errorf("convert: unexpected type for Bool, got type %T", data)
}

func Values(data interface{}) ([]interface{}, error) {
	switch data := data.(type) {
	case []interface{}:
		return data, nil
    case [][]string:
        ret := make([]interface{}, len(data))
        for ind := 0; ind < len(data); ind++ {
            ret[ind] = data[ind]
        }
        return ret, nil
	case nil:
		return nil, ErrNil
	}
	return nil, fmt.Errorf("convert: unexpected type for Values, got type %T", data)
}

func Strings(data interface{}) ([]string, error) {
	switch data := data.(type) {
    case []string:
        return data, nil
	case []interface{}:
		result := make([]string, len(data))
		for i, val := range data {
			result[i], _ = String(val)  //nil will be ""
		}
		return result, nil
	case nil:
		return nil, ErrNil
	}
	return nil, fmt.Errorf("convert: unexpected type for Strings, got type %T", data)
}

func StringMap(data interface{}) (map[string]string, error) {
    switch data := data.(type) {
    case map[string]string:
        return data, nil
    case map[string]interface{}:
        ret := make(map[string]string)
        for key, value := range data {
            ret[key], _ = String(value)
        }
        return ret, nil
    case nil:
        return nil, ErrNil
    }
    return nil, fmt.Errorf("convert: unexpected type for StringMap, got type %T", data)
}
