package utils

import (
	"fmt"
	"strconv"
)

func CheckNumberAndConvertToInt(arg interface{}) (int, error) {
	if arg == nil {
		return int(0), nil
	}
	switch arg.(type) {
	case int:
		return arg.(int), nil
	case int32:
		return int(arg.(int32)), nil
	case int64:
		return int(arg.(int64)), nil
	case int16:
		return int(arg.(int16)), nil
	case int8:
		return int(arg.(int8)), nil
	case float32:
		return int(arg.(float32)), nil
	case float64:
		return int(arg.(float64)), nil
	case string:
		num, err := strconv.Atoi(arg.(string))
		if err != nil {
			return 0, err
		}
		return num, nil
	default:
		return 0, fmt.Errorf("%+v can not convert to int", arg)

	}
}

func CheckNumberAndConvertTofloat64(arg interface{}) (float64, error) {
	if arg == nil {
		return float64(0), nil
	}
	switch arg.(type) {
	case int:
		return float64(arg.(int)), nil
	case int32:
		return float64(arg.(int32)), nil
	case int64:
		return float64(arg.(int64)), nil
	case int16:
		return float64(arg.(int16)), nil
	case int8:
		return float64(arg.(int8)), nil
	case float32:
		return float64(arg.(float32)), nil
	case float64:
		return arg.(float64), nil
	case string:
		num, err := strconv.ParseFloat(arg.(string), 16)
		if err != nil {
			return 0.0, err
		}
		return num, nil
	default:
		return 0.0, fmt.Errorf("%+v can not convert to float64", arg)

	}
}
