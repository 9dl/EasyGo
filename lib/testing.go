package easygo

import "fmt"

func TriggerFakeError() error {
	return fmt.Errorf("This is a fake error")
}
