package easygo

import (
	helpers "easygo/helpers"
	"errors"
	"os"
	"os/exec"
	"runtime"
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
