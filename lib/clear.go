package easygo

import (
	"errors"
	"os"
	"os/exec"
	"runtime"

	helpers "github.com/9dl/EasyGo/helpers"
)

func ConsoleClear() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("clear")
	} else {
		helpers.HandleError(errors.New("helpers.HandleError(err)"))

	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

/*
func main() {
	ConsoleClear()
}
*/
