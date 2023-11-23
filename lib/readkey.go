package easygo

import (
	"bufio"
	"os"

	helpers "github.com/9dl/EasyGo/helpers"
)

func ReadKey() {
	reader := bufio.NewReader(os.Stdin)
	_, _, err := reader.ReadRune()
	if err != nil {
		helpers.HandleError(err)
	}
}

func GetKey() (rune, error) {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		helpers.HandleError(err)
	}
	return char, nil
}

/*
func main() {
	ReadKey()
}

func main() {
	char, err := GetKey()
	if err != nil {
		helpers.HandleError(err)
	} else {
		fmt.Printf("You pressed: %c\n", char)
	}
}
*/
