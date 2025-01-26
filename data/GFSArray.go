package data

import (
	"fmt"
	"strings"
)

type GFSArray struct {
	DataHolder []*GFSDataWrapper
}

func (s *GFSArray) String() string {
	return s.Dump(0)
}

func (s *GFSArray) Dump(indents int) string {
	dump := fmt.Sprintf("(%s):", strings.ToLower(GFS_ARRAY.Name))
	for _, value := range s.DataHolder {
		dump += "\n" + value.Dump("", indents+1)
	}
	return dump
}

func (s *GFSArray) Size() int {
	return len(s.DataHolder)
}

func (s *GFSArray) Get(index int) *GFSDataWrapper {
	if index < 0 || index >= len(s.DataHolder) {
		return nil
	}
	return s.DataHolder[index]
}

func (s *GFSArray) Add(value interface{}) {
	s.AddObject(GFSDataWrapper{
		Data:   value,
		TypeID: GFSDataTypeFromVar(value),
	})
}

func (s *GFSArray) AddObject(wrapper GFSDataWrapper) {
	s.DataHolder = append(s.DataHolder, &wrapper)
}

func (s *GFSArray) RemoveObject(wrapper *GFSDataWrapper) {
	newDataHolder := make([]*GFSDataWrapper, 0, len(s.DataHolder))
	for _, value := range s.DataHolder {
		if value != wrapper {
			newDataHolder = append(newDataHolder, value)
		}
	}
	s.DataHolder = newDataHolder
}

func MakeGFSArray() *GFSArray {
	return &GFSArray{}
}
