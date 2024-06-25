package gofox2x

import "fmt"

type sfsTypeID byte

const (
	SFSNull        sfsTypeID = 0
	SFSBool        sfsTypeID = 1
	SFSByte        sfsTypeID = 2
	SFSShort       sfsTypeID = 3
	SFSInt         sfsTypeID = 4
	SFSLong        sfsTypeID = 5
	SFSFloat       sfsTypeID = 6
	SFSDouble      sfsTypeID = 7
	SFSString      sfsTypeID = 8
	SFSBoolArray   sfsTypeID = 9
	SFSByteArray   sfsTypeID = 10
	SFSShortArray  sfsTypeID = 11
	SFSIntArray    sfsTypeID = 12
	SFSLongArray   sfsTypeID = 13
	SFSFloatArray  sfsTypeID = 14
	SFSDoubleArray sfsTypeID = 15
	SFSStringArray sfsTypeID = 16
	SFSArray       sfsTypeID = 17
	SFSObject      sfsTypeID = 18
	SFSClass       sfsTypeID = 19
	SFSText        sfsTypeID = 20
)

func (s sfsTypeID) String() string {
	switch s {
	case SFSNull:
		return "null"
	case SFSBool:
		return "bool"
	case SFSByte:
		return "byte"
	case SFSShort:
		return "short"
	case SFSInt:
		return "int"
	case SFSLong:
		return "long"
	case SFSFloat:
		return "float"
	case SFSDouble:
		return "double"
	case SFSString:
		return "string"
	case SFSBoolArray:
		return "bool_array"
	case SFSByteArray:
		return "byte_array"
	case SFSShortArray:
		return "short_array"
	case SFSIntArray:
		return "int_array"
	case SFSLongArray:
		return "long_array"
	case SFSFloatArray:
		return "float_array"
	case SFSDoubleArray:
		return "double_array"
	case SFSStringArray:
		return "string_array"
	case SFSArray:
		return "sfs_array"
	case SFSObject:
		return "sfs_object"
	case SFSClass:
		return "class"
	case SFSText:
		return "text"
	default:
		return "unknown"
	}
}

type sfsType struct {
	typeID sfsTypeID
	name   string
	data   interface{}
}

func (st sfsType) String() string {
	return fmt.Sprintf(
		"(%s) %s: %v",
		st.typeID,
		st.name,
		st.data,
	)
}

func makeSFSType(name string, data interface{}) sfsType {
	switch data.(type) {
	case bool:
		return sfsType{typeID: SFSBool, name: name, data: data}
	case int8:
		return sfsType{typeID: SFSByte, name: name, data: data}
	case int16:
		return sfsType{typeID: SFSShort, name: name, data: data}
	case int32:
		return sfsType{typeID: SFSInt, name: name, data: data}
	case int64:
		return sfsType{typeID: SFSLong, name: name, data: data}
	case float32:
		return sfsType{typeID: SFSFloat, name: name, data: data}
	case float64:
		return sfsType{typeID: SFSDouble, name: name, data: data}
	case string:
		return sfsType{typeID: SFSString, name: name, data: data}
	case []bool:
		return sfsType{typeID: SFSBoolArray, name: name, data: data}
	case []int8:
		return sfsType{typeID: SFSByteArray, name: name, data: data}
	case []int16:
		return sfsType{typeID: SFSShortArray, name: name, data: data}
	case []int32:
		return sfsType{typeID: SFSIntArray, name: name, data: data}
	case []int64:
		return sfsType{typeID: SFSLongArray, name: name, data: data}
	case []float32:
		return sfsType{typeID: SFSFloatArray, name: name, data: data}
	case []float64:
		return sfsType{typeID: SFSDoubleArray, name: name, data: data}
	case []string:
		return sfsType{typeID: SFSStringArray, name: name, data: data}
	case struct{}:
		return sfsType{typeID: SFSClass, name: name, data: data}
	default:
		return sfsType{typeID: SFSNull, name: name, data: data}
	}
}
