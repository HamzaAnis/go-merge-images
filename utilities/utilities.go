package utilities

import (
	"fmt"
	"io/ioutil"
)

func GetFiles(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
	return nil
}
