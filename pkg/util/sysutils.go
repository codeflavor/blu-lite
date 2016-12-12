package util

import (
	"os"
)

// CheckCfgDir checks if the directory where the configuration should be found,
// exists
func CheckCfgDir(dir string) error {
	_, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0700)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}
