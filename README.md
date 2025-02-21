# GoFox2X
GoFox2X is a partial port of the [SmartFoxServer2X client](https://www.smartfoxserver.com/products/sfs2x)

## Links:

[GFS Types](#types)

### Tutorial
[Making an Array](#making-an-array)

### Legal
[License](https://github.com/iestyn129/GoFox2X/blob/main/LICENSE)

Making an Array
-
```go
package gofox2x

import (
	"fmt"
	"iestyn129.dev/GoFox2X/data"
	"testing"
)

func main() {
	array := data.MakeGFSArray()
	array.Add(123)
	array.Add("iestyn129")
	array2 := data.MakeGFSArray()
	array2.Add(321)
	var float float32 = 123.321
	array2.Add(float)
	array.Add(array2)
	fmt.Println(array.String())
}
```

This code snippet creates 2 GFSArrays `array` & `array2` and adds elements to them, then adds `array2` to `array` and prints the array.

Types
-
Here is a comprehensive table of all GFS Data Types, IDs and descriptions:

| Name | TypeID | Description |
| --- | --- | --- |
| Null | 0 | TBA |
| Bool | 1 | A true or false boolean |
| Byte | 2 | 8-bit signed integer |
| Short | 3 | 16-bit signed integer |
| Int | 4 | 32-bit signed integer |
| Long | 5 | 64-bit signed integer |
| Float | 6 | Single-precision 32-bit floating point |
| Double | 7 | Double-precision 64-bit floating point |
| Utf-String | 8 | TBA |
| Bool Array | 9 | TBA |
| Byte Array | 10 | TBA |
| Short Array | 11 | TBA |
| Int Array | 12 | TBA |
| Long Array | 13 | TBA |
| Float Array | 14 | TBA |
| Double Array | 15 | TBA |
| Utf-String Array | 16 | TBA |
| GFSArray | 17 | TBA |
| GFSObject | 18 | TBA |
| Class | 19 | TBA |
| Text | 20 | Stores any type of Text Data |
