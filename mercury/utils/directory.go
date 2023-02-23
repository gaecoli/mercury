package utils

import "os"

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true, nil
		}
		return false, nil
	}
	return true, err
}
