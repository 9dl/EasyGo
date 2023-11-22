package easygo

import (
	helpers "easygo/helpers"

	"github.com/gen2brain/dlgs"
)

func ShowErrorMessage(title, message string) {
	_, err := dlgs.Error(title, message)
	if err != nil {
		helpers.HandleError(err)
	}
}

func ShowInfoMessage(title, message string) {
	_, err := dlgs.Info(title, message)
	if err != nil {
		helpers.HandleError(err)
	}
}

func ShowWarningMessage(title, message string) {
	_, err := dlgs.Warning(title, message)
	if err != nil {
		helpers.HandleError(err)
	}
}
