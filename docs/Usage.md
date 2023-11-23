### `getUsage()`

The `getUsage` function retrieves CPU and memory usage information based on the operating system.

**Usage:**

```go
package main

import (
	"fmt"
	"github.com/9dl/EasyGo"
)

func main() {
	cpuUsage, memoryUsage, err := easygo.GetUsage()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)
	fmt.Printf("Memory Usage: %.2f%%\n", memoryUsage)
}
```

[Next Page](/Title.md)
