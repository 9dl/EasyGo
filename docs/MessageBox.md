### `ShowErrorMessage(title, message string)`

The `ShowErrorMessage` function displays an error message dialog.

**Parameters:**

- `title` (string): The title of the error message dialog.
- `message` (string): The content or details of the error message.

**Usage:**

```go
package main

import (
	"github.com/9dl/EasyGo"
)

func main() {
	easygo.ShowErrorMessage("Error", "An unexpected error occurred.")
}
```

---

### `ShowInfoMessage(title, message string)`

The `ShowInfoMessage` function displays an information message dialog.

**Parameters:**

- `title` (string): The title of the information message dialog.
- `message` (string): The content or details of the information message.

**Usage:**

```go
package main

import (
	"github.com/9dl/EasyGo"
)

func main() {
	easygo.ShowInfoMessage("Information", "This is an informative message.")
}
```

---

### `ShowWarningMessage(title, message string)`

The `ShowWarningMessage` function displays a warning message dialog.

**Parameters:**

- `title` (string): The title of the warning message dialog.
- `message` (string): The content or details of the warning message.

**Usage:**

```go
package main

import (
	"github.com/9dl/EasyGo"
)

func main() {
	easygo.ShowWarningMessage("Warning", "Proceed with caution!")
}
```

[Next Page](/docs/ReadKey.md)
