package reflection

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/volatiletech/null"
)

func Set(value reflect.Value, input interface{}) error {
	typ := value.Type()

	switch typ.Kind() {
	case reflect.Int:
		return setIntField(value, input, 0)
	case reflect.Int8:
		return setIntField(value, input, 8)
	case reflect.Int16:
		return setIntField(value, input, 16)
	case reflect.Int32:
		return setIntField(value, input, 32)
	case reflect.Int64:
		return setIntField(value, input, 64)
	case reflect.Uint:
		return setUintField(value, input, 0)
	case reflect.Uint8:
		return setUintField(value, input, 8)
	case reflect.Uint16:
		return setUintField(value, input, 16)
	case reflect.Uint32:
		return setUintField(value, input, 32)
	case reflect.Uint64:
		return setUintField(value, input, 64)
	case reflect.Bool:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.String:
	case reflect.Ptr:
		// special
	default:
		switch typ.String() {
		case "null.Int":
		case "null.Int8":
		case "null.Int16":
		case "null.Int32":
		case "null.Int64":
		case "null.Uint":
		case "null.Uint8":
		case "null.Uint16":
		case "null.Uint32":
		case "null.Uint64":
		case "null.Bool":
		case "null.Float32":
		case "null.Float64":
		case "null.String":
		}
	}

	return nil
}

func getIntValue(value interface{}, bitSize int) (int64, error) {
	var intValue int64
	switch value.(type) {
	case int:
		intValue = int64(value.(int))
	case int8:
		intValue = int64(value.(int8))
	case int16:
		intValue = int64(value.(int16))
	case int32:
		intValue = int64(value.(int32))
	case int64:
		intValue = value.(int64)
	case string:
		if value == "" {
			value = "0"
		}
		var err error
		intValue, err = strconv.ParseInt(value.(string), 10, bitSize)
		if err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("conversion between int type and %s failed", reflect.TypeOf(value).String())
	}

	return intValue, nil
}

func setIntField(field reflect.Value, value interface{}, bitSize int) error {
	intValue, err := getIntValue(value, bitSize)
	if err != nil {
		return err
	}

	field.SetInt(intValue)
	return nil
}

func getUintValue(value interface{}, bitSize int) (uint64, error) {
	var uintValue uint64
	switch value.(type) {
	case uint:
		uintValue = uint64(value.(uint))
	case uint8:
		uintValue = uint64(value.(uint8))
	case uint16:
		uintValue = uint64(value.(uint16))
	case uint32:
		uintValue = uint64(value.(uint32))
	case uint64:
		uintValue = value.(uint64)
	case string:
		if value == "" {
			value = "0"
		}
		var err error
		uintValue, err = strconv.ParseUint(value.(string), 10, bitSize)
		if err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("conversion between uint type and %s failed", reflect.TypeOf(value).String())
	}

	return uintValue, nil
}

func setUintField(field reflect.Value, value interface{}, bitSize int) error {
	uintValue, err := getUintValue(value, bitSize)
	if err != nil {
		return err
	}

	field.SetUint(uintValue)
	return nil
}

func setBoolField(field reflect.Value, value interface{}) error {
	var boolValue = false
	switch value.(type) {
	case string:
		var err error
		boolValue, err = strconv.ParseBool(value.(string))
		if err != nil {
			return err
		}
	case int, int8, int16, int32, int64:
		intValue, err := getIntValue(value, 64)
		if err != nil {
			return err
		}
		switch intValue {
		case 0:
			boolValue = false
		case 1:
			boolValue = true
		default:
			return errors.New("")
		}
	case uint, uint8, uint16, uint32, uint64:

	}

	field.SetBool(boolValue)
	return nil
}

func setNullStringField(value string, field reflect.Value) error {
	if value == "" {
		return nil
	}

	field.Set(reflect.ValueOf(null.StringFrom(value)))
	return nil
}
