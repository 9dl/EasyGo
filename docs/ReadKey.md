### `ReadKey()`

The `ReadKey` function reads a single key press from the standard input.

**Usage:**

```go
package main

import (
	"github.com/9dl/EasyGo"
)

func main() {
	easygo.ReadKey()
}
```

---

### `GetKey() (rune, error)`

The `GetKey` function reads a single key press from the standard input and returns the pressed key as a `rune`.

**Usage:**

```go
package main

import (
	"fmt"
	"github.com/9dl/EasyGo"
)

func main() {
	char, err := easygo.GetKey()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("You pressed: %c\n", char)
}
```

[Next Page](/docs/Strings.md)
