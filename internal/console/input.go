package console

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func GetString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.Replace(text, "\n", "", -1)
	if text == "" {
		return text, errors.New("empty input string")
	}
	return text, nil
}

func GetFile() (*os.File, error) {
	fmt.Println("Input file name:")
	response, err := GetString()
	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(response, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, err
	}
	return file, nil
}
