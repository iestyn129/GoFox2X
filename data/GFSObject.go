package data

import (
	"fmt"
	"strings"
)

type GFSObject struct {
	DataHolder map[string]*GFSDataWrapper
}

func (o *GFSObject) Size() int {
	return len(o.DataHolder)
}

func (o *GFSObject) Dump(indents int) string {
	dump := fmt.Sprintf("(%s):", strings.ToLower(GFS_ARRAY.Name))
	for key, value := range o.DataHolder {
		dump += "\n" + value.Dump(key, indents)
	}
	return dump
}

func (o *GFSObject) ContainsKey(key string) bool {
	_, exists := o.DataHolder[key]
	return exists
}

func (o *GFSObject) RemoveKey(key string) {
	delete(o.DataHolder, key)
}

func (o *GFSObject) Get(key string) *GFSDataWrapper {
	return o.DataHolder[key]
}
