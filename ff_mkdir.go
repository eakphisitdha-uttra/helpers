package helpers

import (
	"fmt"
	"os"
)

func Mkdir(path string, dirMode os.FileMode) error {
	err := os.MkdirAll(path, dirMode)
	if err != nil {
		return fmt.Errorf("%s: creating directory: %v", path, err)
	}
	return nil
}
