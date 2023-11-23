### `ConsoleTitle(title string)`

The `ConsoleTitle` function sets the title of the console window. It uses ANSI escape codes on Linux systems and the Windows API on Windows systems.

**Parameters:**

- `title` (string): The title to be set for the console window.

**Usage:**

```go
package main

import (
	"github.com/9dl/EasyGo"
)

func main() {
	easygo.ConsoleTitle("Hello World")
}
```

[Next Page](/docs/MessageBox.md)
