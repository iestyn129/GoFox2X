package data

import "strings"

type GFSDataType struct {
	TypeID byte
	Name   string
}

var (
	NULL             = GFSDataType{0, "null"}
	BOOL             = GFSDataType{1, "bool"}
	BYTE             = GFSDataType{2, "byte"}
	SHORT            = GFSDataType{3, "short"}
	INT              = GFSDataType{4, "int"}
	LONG             = GFSDataType{5, "long"}
	FLOAT            = GFSDataType{6, "float"}
	DOUBLE           = GFSDataType{7, "double"}
	UTF_STRING       = GFSDataType{8, "utf_string"}
	BOOL_ARRAY       = GFSDataType{9, "bool_array"}
	BYTE_ARRAY       = GFSDataType{10, "byte_array"}
	SHORT_ARRAY      = GFSDataType{11, "short_array"}
	INT_ARRAY        = GFSDataType{12, "int_array"}
	LONG_ARRAY       = GFSDataType{13, "long_array"}
	FLOAT_ARRAY      = GFSDataType{14, "float_array"}
	DOUBLE_ARRAY     = GFSDataType{15, "double_array"}
	UTF_STRING_ARRAY = GFSDataType{16, "utf_string_array"}
	GFS_ARRAY        = GFSDataType{17, "gfs_array"}
	GFS_OBJECT       = GFSDataType{18, "gfs_object"}
	CLASS            = GFSDataType{19, "class"}
	TEXT             = GFSDataType{20, "text"}
)

var GFSDataTypes = []GFSDataType{
	NULL, BOOL, BYTE, SHORT, INT, LONG, FLOAT, DOUBLE, UTF_STRING,
	BOOL_ARRAY, BYTE_ARRAY, SHORT_ARRAY, INT_ARRAY, LONG_ARRAY, FLOAT_ARRAY, DOUBLE_ARRAY, UTF_STRING_ARRAY,
	GFS_ARRAY, GFS_OBJECT, CLASS, TEXT,
}

func (d GFSDataType) String() string {
	return strings.ToLower(d.Name)
}

func GFSDataTypeFromVar(value interface{}) GFSDataType {
	switch value.(type) {
	case nil:
		return NULL
	case bool:
		return BOOL
	case byte:
		return BYTE
	case int16:
		return SHORT
	case int:
		return INT
	case int64:
		return LONG
	case float32:
		return FLOAT
	case float64:
		return DOUBLE
	case string:
		return UTF_STRING
	case []bool:
		return BOOL_ARRAY
	case []byte:
		return BYTE_ARRAY
	case []int16:
		return SHORT_ARRAY
	case []int:
		return INT_ARRAY
	case []int64:
		return LONG_ARRAY
	case []float32:
		return FLOAT_ARRAY
	case []float64:
		return DOUBLE_ARRAY
	case []string:
		return UTF_STRING_ARRAY
	case *GFSArray:
		return GFS_ARRAY
	// TODO:
	//case GFSObject:
	//	return
	default:
		return NULL
	}
}

func GFSDataTypeFromTypeID(typeID byte) *GFSDataType {
	for _, dt := range GFSDataTypes {
		if dt.TypeID == typeID {
			return &dt
		}
	}
	return nil
}
