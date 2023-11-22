package helpers

import (
	"fmt"
	"runtime"

	"github.com/logrusorgru/aurora"
)

func handleError(err error) {
	if err != nil {
		pc, filename, line, _ := runtime.Caller(1)

		os := runtime.GOOS
		goVersion := runtime.Version()

		fmt.Printf("╭─ [%s] in %s[%s:%d] %s\n", aurora.Red(aurora.Bold("error")), aurora.Blue(runtime.FuncForPC(pc).Name()), aurora.Green(filename), aurora.Cyan(line), aurora.Cyan("↙"))
		fmt.Printf("│   %v\n", err)
		fmt.Printf("│\n")
		fmt.Printf("│  %s Something went wrong! \n", aurora.Cyan("↑"))
		fmt.Printf("│  %s Please check the error message above for details.\n", aurora.Cyan("→"))
		fmt.Printf("│  %s User OS: %s\n", aurora.Cyan("→"), os)
		fmt.Printf("│  %s Go Version: %s\n", aurora.Cyan("→"), goVersion)
		fmt.Println("╰───────────────────────────────────")
		os.Exit(1)
	}
}
