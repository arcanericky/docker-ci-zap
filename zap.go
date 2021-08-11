// +build windows

package dockercizap

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Microsoft/hcsshim"
)

var destroyer func(hcsshim.DriverInfo, string) error = hcsshim.DestroyLayer
var folderChecker func(string) bool = folderExists

func folderExists(folder string) bool {
	if fi, err := os.Stat(folder); err != nil {
		return false
	} else if !fi.IsDir() {
		return false
	}

	return true
}

func destroyLayer(folder string) error {
	location, folderName := filepath.Split(folder)

	return destroyer(hcsshim.DriverInfo{
		HomeDir: location,
		Flavour: 0,
	}, folderName)
}

func Zap(folder string) error {
	if !folderChecker(folder) {
		return errors.New("folder does not exist")
	}

	if err := destroyLayer(folder); err != nil {
		return err
	}

	return nil
}
