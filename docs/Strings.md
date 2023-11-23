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

---

### `ReplaceString(input string, oldSubstr string, newSubstr string) string`

The `ReplaceString` function replaces all occurrences of `oldSubstr` with `newSubstr` in the given input string.

**Usage:**

```go
package main

import (
	"fmt"
	"github.com/9dl/EasyGo"
)

func main() {
	input := "Hello, World!"
	newString := easygo.ReplaceString(input, "Hello", "Hi")
	fmt.Println(newString)
}
```

### `GetStringLength(input string) int`

The `GetStringLength` function returns the length of the input string.

**Usage:**

```go
package main

import (
	"fmt"
	"github.com/9dl/EasyGo"
)

func main() {
	input := "Hello, World!"
	length := easygo.GetStringLength(input)
	fmt.Printf("Length of the string: %d\n", length)
}
```

### `ReverseString(input string) string`

The `ReverseString` function reverses the characters in the input string.

**Usage:**

```go
package main

import (
	"fmt"
	"github.com/9dl/EasyGo"
)

func main() {
	input := "Hello, World!"
	reversed := easygo.ReverseString(input)
	fmt.Println(reversed)
}
```

### `SplitString(input, delimiter string) []string`

The `SplitString` function splits the input string into a slice of substrings using the specified delimiter.

**Usage:**

```go
package main

import (
	"fmt"
	"github.com/9dl/EasyGo"
)

func main() {
	input := "apple,orange,banana"
	result := easygo.SplitString(input, ",")
	fmt.Println(result)
}
```

### `JoinStrings(elements []string, separator string) string`

The `JoinStrings` function joins the elements of a string slice into a single string, using the specified separator.

**Usage:**

```go
package main

import (
	"fmt"
	"github.com/9dl/EasyGo"
)

func main() {
	elements := []string{"apple", "orange", "banana"}
	joined := easygo.JoinStrings(elements, ", ")
	fmt.Println(joined)
}
```

### `ContainsSubstring(input, substring string) bool`

The `ContainsSubstring` function checks if the input string contains the specified substring.

**Usage:**

```go
package main

import (
	"fmt"
	"github.com/9dl/EasyGo"
)

func main() {
	input := "Hello, World!"
	contains := easygo.ContainsSubstring(input, "World")
	fmt.Printf("Contains 'World': %v\n", contains)
}
```

### `ToLowerCase(input string) string`

The `ToLowerCase` function converts the input string to lowercase.

**Usage:**

```go
package main

import (
	"fmt"
	"github.com/9dl/EasyGo"
)

func main() {
	input := "Hello, World!"
	lowercase := easygo.ToLowerCase(input)
	fmt.Println(lowercase)
}
```

### `ToUpperCase(input string) string`

The `ToUpperCase` function converts the input string to uppercase.

**Usage:**

```go
package main

import (
	"fmt"
	"github.com/9dl/EasyGo"
)

func main() {
	input := "Hello, World!"
	uppercase := easygo.ToUpperCase(input)
	fmt.Println(uppercase)
}
```

### `TrimSpaces(input string) string`

The `TrimSpaces` function removes leading and trailing whitespaces from the input string.

**Usage:**

```go
package main

import (
	"fmt"
	"github.com/9dl/EasyGo"
)

func main() {
	input := "   Hello, World!   "
	trimmed := easygo.TrimSpaces(input)
	fmt.Println(trimmed)
}
```

[Next Page](/docs/File.md)
