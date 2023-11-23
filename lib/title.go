package easygo

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	helpers "github.com/9dl/EasyGo/helpers"
)

func ConsoleTitle(title string) {
	if _, ok := os.LookupEnv("TERM"); !ok && runtime.GOOS != "windows" {
		helpers.HandleError(errors.New("Terminal does not support ANSI escape codes"))
	}

	fmt.Printf("\033]0;%s\007", title)
}

/*
func main() {
	ConsoleTitle("Hello World")
}
*/
