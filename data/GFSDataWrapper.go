package data

import (
	"fmt"
	"strings"
)

type GFSDataWrapper struct {
	Data   interface{}
	TypeID GFSDataType
}

func (w *GFSDataWrapper) Dump(key string, indents int) string {
	indent := strings.Repeat("\t", indents)
	if key != "" {
		key = " " + key
	}

	typeToString := func(typeName string, data interface{}) string {
		return fmt.Sprintf("%s(%s)%s: %v", indent, strings.ToLower(typeName), key, data)
	}

	switch w.TypeID {
	case BOOL_ARRAY:
		return typeToString(w.TypeID.Name, joinArray(w.Data.([]bool)))
	case BYTE_ARRAY:
		return typeToString(w.TypeID.Name, joinArray(w.Data.([]byte)))
	case SHORT_ARRAY:
		return typeToString(w.TypeID.Name, joinArray(w.Data.([]int16)))
	case INT_ARRAY:
		return typeToString(w.TypeID.Name, joinArray(w.Data.([]int)))
	case LONG_ARRAY:
		return typeToString(w.TypeID.Name, joinArray(w.Data.([]int64)))
	case FLOAT_ARRAY:
		return typeToString(w.TypeID.Name, joinArray(w.Data.([]float32)))
	case DOUBLE_ARRAY:
		return typeToString(w.TypeID.Name, joinArray(w.Data.([]float64)))
	case UTF_STRING_ARRAY:
		return typeToString(w.TypeID.Name, joinArray(w.Data.([]string)))
	case GFS_ARRAY:
		return fmt.Sprintf("%s%v", indent, w.Data.(*GFSArray).Dump(indents))
	// TODO:
	//case GFS_OBJECT:
	//	return typeToString(w.TypeID.Name, w.Data.(*GFSObject).Dump(indents + 1))
	case NULL:
		return typeToString(w.TypeID.Name, "null")
	default:
		return typeToString(w.TypeID.Name, w.Data)
	}
}

func joinArray[T any](array []T) string {
	str := ""
	for _, v := range array {
		str += fmt.Sprintf(", %v", v)
	}
	return strings.TrimPrefix(str, ", ")
}

func NewGFSDataWrapper(data interface{}) *GFSDataWrapper {
	return &GFSDataWrapper{
		Data:   data,
		TypeID: GFSDataTypeFromVar(data),
	}
}
