package gofox2x

import "fmt"

type sfsTypeID byte

const (
	SFSNULL        sfsTypeID = 0
	SFSBOOL        sfsTypeID = 1
	SFSBYTE        sfsTypeID = 2
	SFSSHORT       sfsTypeID = 3
	SFSINT         sfsTypeID = 4
	SFSLONG        sfsTypeID = 5
	SFSFLOAT       sfsTypeID = 6
	SFSDOUBLE      sfsTypeID = 7
	SFSSTRING      sfsTypeID = 8
	SFSBOOLARRAY   sfsTypeID = 9
	SFSBYTEARRAY   sfsTypeID = 10
	SFSSHORTARRAY  sfsTypeID = 11
	SFSINTARRAY    sfsTypeID = 12
	SFSLONGARRAY   sfsTypeID = 13
	SFSFLOATARRAY  sfsTypeID = 14
	SFSDOUBLEARRAY sfsTypeID = 15
	SFSSTRINGARRAY sfsTypeID = 16
	SFSARRAY       sfsTypeID = 17
	SFSOBJECT      sfsTypeID = 18
	SFSCLASS       sfsTypeID = 19
	SFSTEXT        sfsTypeID = 20
)

func (s sfsTypeID) String() string {
	switch s {
	case SFSNULL:
		return "null"
	case SFSBOOL:
		return "bool"
	case SFSBYTE:
		return "byte"
	case SFSSHORT:
		return "short"
	case SFSINT:
		return "int"
	case SFSLONG:
		return "long"
	case SFSFLOAT:
		return "float"
	case SFSDOUBLE:
		return "double"
	case SFSSTRING:
		return "string"
	case SFSBOOLARRAY:
		return "bool_array"
	case SFSBYTEARRAY:
		return "byte_array"
	case SFSSHORTARRAY:
		return "short_array"
	case SFSINTARRAY:
		return "int_array"
	case SFSLONGARRAY:
		return "long_array"
	case SFSFLOATARRAY:
		return "float_array"
	case SFSDOUBLEARRAY:
		return "double_array"
	case SFSSTRINGARRAY:
		return "string_array"
	case SFSARRAY:
		return "sfs_array"
	case SFSOBJECT:
		return "sfs_object"
	case SFSCLASS:
		return "class"
	case SFSTEXT:
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
		return sfsType{typeID: SFSBOOL, name: name, data: data}
	case int8:
		return sfsType{typeID: SFSBYTE, name: name, data: data}
	case int16:
		return sfsType{typeID: SFSSHORT, name: name, data: data}
	case int32:
		return sfsType{typeID: SFSINT, name: name, data: data}
	case int64:
		return sfsType{typeID: SFSLONG, name: name, data: data}
	case float32:
		return sfsType{typeID: SFSFLOAT, name: name, data: data}
	case float64:
		return sfsType{typeID: SFSDOUBLE, name: name, data: data}
	case string:
		return sfsType{typeID: SFSSTRING, name: name, data: data}
	case []bool:
		return sfsType{typeID: SFSBOOLARRAY, name: name, data: data}
	case []int8:
		return sfsType{typeID: SFSBYTEARRAY, name: name, data: data}
	case []int16:
		return sfsType{typeID: SFSSHORTARRAY, name: name, data: data}
	case []int32:
		return sfsType{typeID: SFSINTARRAY, name: name, data: data}
	case []int64:
		return sfsType{typeID: SFSLONGARRAY, name: name, data: data}
	case []float32:
		return sfsType{typeID: SFSFLOATARRAY, name: name, data: data}
	case []float64:
		return sfsType{typeID: SFSDOUBLEARRAY, name: name, data: data}
	case []string:
		return sfsType{typeID: SFSSTRINGARRAY, name: name, data: data}
	case struct{}:
		return sfsType{typeID: SFSCLASS, name: name, data: data}
	default:
		return sfsType{typeID: SFSNULL, name: name, data: nil}
	}
}
