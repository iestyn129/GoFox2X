package gofox2x

import (
	"bytes"
	"encoding/binary"
)

type SFSObject struct {
	name string
	data map[string]sfsType
}

func NewSFSObject(name string) *SFSObject {
	return &SFSObject{
		name: name,
		data: make(map[string]sfsType),
	}
}

func (s *SFSObject) Encode() ([]byte, error) {
	var buf bytes.Buffer
	var err error

	err = binary.Write(&buf, binary.BigEndian, SFSOBJECT)
	if err != nil {
		return nil, err
	}
	err = binary.Write(&buf, binary.BigEndian, s.Size())
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (s *SFSObject) Size() int16 {
	return int16(len(s.data))
}

func (s *SFSObject) GetKeys() []string {
	keys := make([]string, 0, len(s.data))
	for key := range s.data {
		keys = append(keys, key)
	}
	return keys
}

func (s *SFSObject) Put(key string, value interface{}) {
	switch value.(type) {
	case struct{}:
		s.PutClass(key, value)
	default:
		s.PutNull(key)
	}
}
