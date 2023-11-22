package easygo

import "fmt"

func ConsoleTitle(title string) {
	fmt.Printf("\033]0;%s\007", title)
}

/*
func main() {
	ConsoleTitle("Hello World")
}
*/
