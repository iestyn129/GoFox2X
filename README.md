# GoFox2X
GoFox2X is a partial port of the [SmartFoxServer2X client](https://www.smartfoxserver.com/products/sfs2x)

## Links:

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

This code snippet creates 2 SFSArrays `array` & `array2` and adds elements to them, then adds `array2` to `array` and prints the array.
