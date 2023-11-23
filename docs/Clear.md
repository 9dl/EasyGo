### `ConsoleClear()`

The `ConsoleClear` function clears the console screen, providing a clean slate for your application's output. It uses platform-specific commands to achieve this.

**Usage:**

```go
package main

import (
	"fmt"
	"github.com/9dl/EasyGo"
)

func main() {
	fmt.Println("This is some content on the console.")
	easygo.ConsoleClear()
	fmt.Println("After clearing the console.")
}
```

[Next Page](/docs/Usage.md)
