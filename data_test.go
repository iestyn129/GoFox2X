package gofox2x

import (
	"fmt"
	"iestyn129.dev/GoFox2X/data"
	"testing"
)

func TestGFSArray(t *testing.T) {
	array := data.MakeGFSArray()
	array.Add(123)
	array.Add("iestyn129")
	array.Add(420.69)
	array2 := data.MakeGFSArray()
	array2.Add(321)
	array2.Add("penis")
	var float float32 = 69.420
	array2.Add(float)
	array.Add(array2)
	fmt.Println(array.String())
}
