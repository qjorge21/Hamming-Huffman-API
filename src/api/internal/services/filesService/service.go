package filesservice

import (
	"context"
	"fmt"
	"io/ioutil"
)

func ReadFile(ctx context.Context) (string, error) {
	fileString := ""

	file, err := ioutil.ReadFile("./ejemplo.txt")

	if err != nil {
		fmt.Println(err.Error())
		return fileString, err
	}

	fileString = string(file)

	return fileString, nil
}
