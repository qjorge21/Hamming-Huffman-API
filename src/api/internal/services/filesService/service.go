package filesservice

import (
	"context"
	"io/ioutil"
)

func ReadFile(ctx context.Context){
	file, err := ioutil.ReadFile("../")
}