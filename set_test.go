package reflection

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/volatiletech/null"
)

type testStruct struct {
	IntInt         int   `json:"int_int"`
	IntInt8        int   `json:"int_int_8"`
	IntInt16       int   `json:"int_int_16"`
	IntInt32       int   `json:"int_int_32"`
	IntInt64       int   `json:"int_int_64"`
	Int8Int        int8  `json:"int_8_int"`
	Int16Int       int16 `json:"int_16_int"`
	Int32Int       int32 `json:"int_32_int"`
	Int64Int       int64 `json:"int_64_int"`
	IntString      int   `json:"int_string"`
	IntStringEmpty int   `json:"int_string_empty"`
	IntPtrInt      *int  `json:"int_ptr_int"`

	UintUint   uint   `json:"uint_uint"`
	UintUint8  uint   `json:"uint_uint_8"`
	UintUint16 uint   `json:"uint_uint_16"`
	UintUint32 uint   `json:"uint_uint_32"`
	UintUint64 uint   `json:"uint_uint_64"`
	Uint8Uint  uint8  `json:"uint_8_uint"`
	Uint16Uint uint16 `json:"uint_16_uint"`
	Uint32Uint uint32 `json:"uint_32_uint"`
	Uint64Uint uint64 `json:"uint_64_uint"`
	UintString uint   `json:"uint_string"`

	BoolBool   bool `json:"bool_bool"`
	BoolInt    bool `json:"bool_int"`
	BoolUint   bool `json:"bool_uint"`
	BoolString bool `json:"bool_string"`

	Float32Float32 float32 `json:"float_32_float_32"`
	Float32Float64 float32 `json:"float_32_float_64"`
	Float64Float32 float64 `json:"float_64_float_32"`
	Float64Float64 float64 `json:"float_64_float_64"`
	Float64Int     float64 `json:"float_64_int"`
	Float64Uint    float64 `json:"float_64_uint"`
	Float64String  float64 `json:"float_64_string"`

	StringString string `json:"string_string"`
	StringInt    string `json:"string_int"`
	StringUint   string `json:"string_uint"`
	StringBool   string `json:"string_bool"`
	StringFloat  string `json:"string_float"`

	NullIntInt          null.Int     `json:"null_int_int"`
	NullIntString       null.Int     `json:"null_int_string"`
	NullIntPtrString    *null.Int    `json:"null_int_ptr_string"`
	NullStringInt       null.String  `json:"null_string_int"`
	NullStringPtrString *null.String `json:"null_string_ptr_string"`
}

func TestSet(t *testing.T) {
	dataset := dataSet()

	var ts = &testStruct{}

	rv := reflect.Indirect(reflect.ValueOf(ts))
	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		rvField := rv.Field(i)
		rtField := rt.Field(i)

		if tag, ok := rtField.Tag.Lookup("json"); ok {
			if data, ok := dataset[tag]; ok {
				if err := Set(rvField, data); err != nil {
					t.Error(err)
				}
			}

		}
	}

	assert.Equal(t, 1000, ts.IntInt)
	assert.Equal(t, 20, ts.IntInt8)
	assert.Equal(t, 200, ts.IntInt16)
	assert.Equal(t, 2000, ts.IntInt32)
	assert.Equal(t, 20000, ts.IntInt64)
	assert.Equal(t, int8(10), ts.Int8Int)
	assert.Equal(t, int16(1000), ts.Int16Int)
	assert.Equal(t, int32(1000), ts.Int32Int)
	assert.Equal(t, int64(1000), ts.Int64Int)
	assert.Equal(t, 1000, ts.IntString)
	assert.Equal(t, 0, ts.IntStringEmpty)
	assert.Equal(t, 123, *ts.IntPtrInt)

	assert.Equal(t, uint(1000), ts.UintUint)
	assert.Equal(t, uint(10), ts.UintUint8)
	assert.Equal(t, uint(100), ts.UintUint16)
	assert.Equal(t, uint(1000), ts.UintUint32)
	assert.Equal(t, uint(10000), ts.UintUint64)
	assert.Equal(t, uint8(10), ts.Uint8Uint)
	assert.Equal(t, uint16(1000), ts.Uint16Uint)
	assert.Equal(t, uint32(1000), ts.Uint32Uint)
	assert.Equal(t, uint64(1000), ts.Uint64Uint)
	assert.Equal(t, uint(1000), ts.UintString)

	assert.Equal(t, true, ts.BoolBool)
	assert.Equal(t, true, ts.BoolInt)
	assert.Equal(t, true, ts.BoolString)
	assert.Equal(t, true, ts.BoolUint)

	assert.Equal(t, float32(123.2), ts.Float32Float32)
	assert.Equal(t, float32(1234.2), ts.Float32Float64)
	assert.Equal(t, 123.19999694824219, ts.Float64Float32) // @TODO .... fix this shit
	assert.Equal(t, 1234.2, ts.Float64Float64)
	assert.Equal(t, float64(123), ts.Float64Int)
	assert.Equal(t, float64(123), ts.Float64Uint)
	assert.Equal(t, 123.3, ts.Float64String)

	assert.Equal(t, "test", ts.StringString)
	assert.Equal(t, "123", ts.StringInt)
	assert.Equal(t, "123", ts.StringUint)
	assert.Equal(t, "true", ts.StringBool)
	assert.Equal(t, "1234.2", ts.StringFloat)

	assert.Equal(t, 12, ts.NullIntInt.Int)
	assert.Equal(t, 12, ts.NullIntString.Int)
	assert.Equal(t, 12, ts.NullIntPtrString.Int)
	assert.Equal(t, "1234", ts.NullStringInt.String)
	assert.Equal(t, "test", ts.NullStringPtrString.String)
get rid of the huge
	//res, _ := json.Marshal(ts)
	//fmt.Println(string(res))
}

func dataSet() map[string]interface{} {
	return map[string]interface{}{
		"int_int":          1000,
		"int_int_8":        int8(20),
		"int_int_16":       int16(200),
		"int_int_32":       int32(2000),
		"int_int_64":       int64(20000),
		"int_8_int":        10,
		"int_16_int":       1000,
		"int_32_int":       1000,
		"int_64_int":       1000,
		"int_string":       "1000",
		"int_string_empty": "",
		"int_ptr_int":      123,

		"uint_uint":    uint(1000),
		"uint_uint_8":  uint8(10),
		"uint_uint_16": uint16(100),
		"uint_uint_32": uint32(1000),
		"uint_uint_64": uint64(10000),
		"uint_8_uint":  uint(10),
		"uint_16_uint": uint(1000),
		"uint_32_uint": uint(1000),
		"uint_64_uint": uint(1000),
		"uint_string":  "1000",

		"bool_bool":   true,
		"bool_int":    1,
		"bool_uint":   uint(1),
		"bool_string": "true",

		"float_32_float_32": float32(123.2),
		"float_32_float_64": 1234.2,
		"float_64_float_32": float32(123.2),
		"float_64_float_64": 1234.2,
		"float_64_int":      123,
		"float_64_uint":     uint(123),
		"float_64_string":   "123.3",

		"string_string": "test",
		"string_int":    123,
		"string_uint":   uint(123),
		"string_bool":   true,
		"string_float":  1234.2,

		"null_int_int":           12,
		"null_int_string":        "12",
		"null_int_ptr_string":    "12",
		"null_string_int":        1234,
		"null_string_ptr_string": "test",
	}
}

func Test_setStructField(t *testing.T) {
	type args struct {
		field reflect.Value
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with null typed struct",
			args: args{
				field: reflect.ValueOf(&null.Uint64{}),
				value: "13",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := setStructField(tt.args.field, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("setStructField() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				if value, ok := tt.args.field.Interface().(*null.Uint64); ok {
					if value.Uint64 != 13 {
						t.Errorf("Value should be 13, instead of %d", value.Uint64)
					}
				} else {
					t.Error("Invalid data")
				}
			}
		})
	}
}

func BenchmarkSet(b *testing.B) {
	dataset := dataSet()

	var ts = &testStruct{}

	for i := 0; i < b.N; i++ {
		rv := reflect.Indirect(reflect.ValueOf(ts))
		rt := rv.Type()

		for i := 0; i < rt.NumField(); i++ {
			rvField := rv.Field(i)
			rtField := rt.Field(i)

			if tag, ok := rtField.Tag.Lookup("json"); ok {
				if data, ok := dataset[tag]; ok {
					if err := Set(rvField, data); err != nil {
						b.Error(err)
					}
				}

			}
		}
	}
}
