package utilz

import (
	"errors"
	"os"
)

// PathExists 文件目录是否存在
// @param: path string
// @return: bool, error
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
