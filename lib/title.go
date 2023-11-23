package easygo

import (
	"fmt"
	"os"
	"runtime"
)

func ConsoleTitle(title string) {
	if _, ok := os.LookupEnv("TERM"); !ok && runtime.GOOS != "windows" {
		fmt.Println("Terminal does not support ANSI escape codes. Unable to set console title.")
		return
	}

	fmt.Printf("\033]0;%s\007", title)
}

/*
func main() {
	ConsoleTitle("Hello World")
}
*/
