### `ReadLines(filePath string) []string`

The `ReadLines` function reads the lines from a file specified by the given file path and returns them as a string slice.

**Parameters:**

- `filePath` (string): The path to the file to be read.

**Usage:**

```go
package main

import (
	"fmt"
	"github.com/9dl/EasyGo"
)

func main() {
	filePath := "example.txt"
	lines := easygo.ReadLines(filePath)
	fmt.Println(lines)
}
```

---

### `WriteLines(filePath string, lines []string) bool`

The `WriteLines` function writes the specified lines to a file specified by the given file path.

**Parameters:**

- `filePath` (string): The path to the file to be written.
- `lines` ([]string): The lines to be written to the file.

**Usage:**

```go
package main

import (
	"github.com/9dl/EasyGo"
)

func main() {
	filePath := "example.txt"
	lines := []string{"Line 1", "Line 2", "Line 3"}
	success := easygo.WriteLines(filePath, lines)
	if success {
		fmt.Println("File written successfully.")
	} else {
		fmt.Println("Error writing the file.")
	}
}
```

---

### `FileExists(filePath string) bool`

The `FileExists` function checks if a file exists at the specified file path.

**Parameters:**

- `filePath` (string): The path to the file to be checked.

**Usage:**

```go
package main

import (
	"fmt"
	"github.com/9dl/EasyGo"
)

func main() {
	filePath := "example.txt"
	exists := easygo.FileExists(filePath)
	fmt.Printf("File exists: %v\n", exists)
}
```

---

### `DeleteFile(filePath string) bool`

The `DeleteFile` function deletes the file specified by the given file path.

**Parameters:**

- `filePath` (string): The path to the file to be deleted.

**Usage:**

```go
package main

import (
	"fmt"
	"github.com/9dl/EasyGo"
)

func main() {
	filePath := "example.txt"
	success := easygo.DeleteFile(filePath)
	if success {
		fmt.Println("File deleted successfully.")
	} else {
		fmt.Println("Error deleting the file.")
	}
}
```

[End](/docs/Installation.md)
