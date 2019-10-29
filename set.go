package reflection

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type capableStruct interface {
	UnmarshalText(text []byte) error
}

func Set(field reflect.Value, input interface{}) error {
	switch field.Kind() {
	case reflect.Int:
		return setIntField(field, input, 0)
	case reflect.Int8:
		return setIntField(field, input, 8)
	case reflect.Int16:
		return setIntField(field, input, 16)
	case reflect.Int32:
		return setIntField(field, input, 32)
	case reflect.Int64:
		return setIntField(field, input, 64)
	case reflect.Uint:
		return setUintField(field, input, 0)
	case reflect.Uint8:
		return setUintField(field, input, 8)
	case reflect.Uint16:
		return setUintField(field, input, 16)
	case reflect.Uint32:
		return setUintField(field, input, 32)
	case reflect.Uint64:
		return setUintField(field, input, 64)
	case reflect.Bool:
		return setBoolField(field, input)
	case reflect.Float32:
		return setFloatField(field, input, 32)
	case reflect.Float64:
		return setFloatField(field, input, 64)
	case reflect.String:
		return setStringField(field, input)
	case reflect.Ptr:
		field.Set(reflect.New(reflect.TypeOf(field.Interface()).Elem()))
		return Set(reflect.Indirect(field), input)
	case reflect.Struct:
		return setStructField(field, input)
	default:

	}

	return nil
}

func getIntValue(input interface{}, bitSize int) (int64, error) {
	var intValue int64
	switch input.(type) {
	case int:
		intValue = int64(input.(int))
	case int8:
		intValue = int64(input.(int8))
	case int16:
		intValue = int64(input.(int16))
	case int32:
		intValue = int64(input.(int32))
	case int64:
		intValue = input.(int64)
	case string:
		if input == "" {
			input = "0"
		}
		var err error
		intValue, err = strconv.ParseInt(input.(string), 10, bitSize)
		if err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("conversion between int type and %s failed", reflect.TypeOf(input).String())
	}

	return intValue, nil
}

func setIntField(field reflect.Value, input interface{}, bitSize int) error {
	intValue, err := getIntValue(input, bitSize)
	if err != nil {
		return err
	}

	field.SetInt(intValue)
	return nil
}

func getUintValue(input interface{}, bitSize int) (uint64, error) {
	var uintValue uint64
	switch input.(type) {
	case uint:
		uintValue = uint64(input.(uint))
	case uint8:
		uintValue = uint64(input.(uint8))
	case uint16:
		uintValue = uint64(input.(uint16))
	case uint32:
		uintValue = uint64(input.(uint32))
	case uint64:
		uintValue = input.(uint64)
	case string:
		if input == "" {
			input = "0"
		}
		var err error
		uintValue, err = strconv.ParseUint(input.(string), 10, bitSize)
		if err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("conversion between uint type and %s failed", reflect.TypeOf(input).String())
	}

	return uintValue, nil
}

func setUintField(field reflect.Value, input interface{}, bitSize int) error {
	uintValue, err := getUintValue(input, bitSize)
	if err != nil {
		return err
	}

	field.SetUint(uintValue)
	return nil
}

func setBoolField(field reflect.Value, input interface{}) error {
	var boolValue = false
	switch input.(type) {
	case bool:
		boolValue = input.(bool)
	case string:
		var err error
		boolValue, err = strconv.ParseBool(input.(string))
		if err != nil {
			return err
		}
	case int, int8, int16, int32, int64:
		intValue, err := getIntValue(input, 64)
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
		uintValue, err := getUintValue(input, 64)
		if err != nil {
			return err
		}
		switch uintValue {
		case 0:
			boolValue = false
		case 1:
			boolValue = true
		default:
			return errors.New("")
		}
	default:
		return fmt.Errorf("conversion between bool type and %s failed", reflect.TypeOf(input).String())
	}

	field.SetBool(boolValue)
	return nil
}

func getStringValue(input interface{}) (string, error) {
	var stringValue string
	switch input.(type) {
	case string:
		stringValue = input.(string)
	case int, int8, int16, int32, int64:
		inputValue, err := getIntValue(input, 64)
		if err != nil {
			return "", err
		}
		stringValue = strconv.FormatInt(inputValue, 10)
	case uint, uint8, uint16, uint32, uint64:
		uintValue, err := getUintValue(input, 64)
		if err != nil {
			return "", err
		}
		stringValue = strconv.FormatUint(uintValue, 10)
	case bool:
		stringValue = strconv.FormatBool(input.(bool))
	case float32, float64:
		floatValue, err := getFloatValue(input, 64)
		if err != nil {
			return "", err
		}
		stringValue = strconv.FormatFloat(floatValue, 'f', -1, 64)
	default:
		return "", fmt.Errorf("conversion between string type and %s failed", reflect.TypeOf(input))
	}

	return stringValue, nil
}

func setStringField(field reflect.Value, input interface{}) error {
	stringValue, err := getStringValue(input)
	if err != nil {
		return err
	}

	field.SetString(stringValue)
	return nil
}

func getFloatValue(input interface{}, bitSize int) (float64, error) {
	var floatValue float64
	switch input.(type) {
	case float32:
		floatValue = float64(input.(float32))
	case float64:
		floatValue = input.(float64)
	case int, int8, int16, int32, int64:
		inputValue, err := getIntValue(input, 64)
		if err != nil {
			return 0, err
		}
		floatValue = float64(inputValue)
	case uint, uint8, uint16, uint32, uint64:
		uintValue, err := getUintValue(input, 64)
		if err != nil {
			return 0, err
		}
		floatValue = float64(uintValue)
	case string:
		if input == "" {
			input = "0.0"
		}
		var err error
		floatValue, err = strconv.ParseFloat(input.(string), bitSize)
		if err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("conversion between float type and %s failed", reflect.TypeOf(input).String())
	}

	return floatValue, nil
}

func setFloatField(field reflect.Value, input interface{}, bitSize int) error {
	floatValue, err := getFloatValue(input, bitSize)
	if err != nil {
		return err
	}

	field.SetFloat(floatValue)
	return nil
}

func setStructField(field reflect.Value, input interface{}) error {
	stringValue, err := getStringValue(input)
	if err != nil {
		return err
	}

	if field.Kind() == reflect.Ptr {
		field = reflect.Indirect(field)
	}

	newField := reflect.New(field.Type())
	if casted, ok := newField.Interface().(capableStruct); ok {
		err := casted.UnmarshalText([]byte(stringValue))
		if err != nil {
			return err
		}
	}

	field.Set(newField.Elem())

	return nil
}
