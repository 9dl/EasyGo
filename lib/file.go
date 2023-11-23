package easygo

import (
	"bufio"
	"os"

	helpers "github.com/9dl/EasyGo/helpers"
)

func ReadLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		helpers.HandleError(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		helpers.HandleError(err)
	}

	return lines
}

func WriteLines(filePath string, lines []string) bool {
	file, err := os.Create(filePath)
	if err != nil {
		helpers.HandleError(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			helpers.HandleError(err)
		}
	}

	err = writer.Flush()
	if err != nil {
		helpers.HandleError(err)
	}

	return true
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func DeleteFile(filePath string) bool {
	err := os.Remove(filePath)
	return err == nil
}
