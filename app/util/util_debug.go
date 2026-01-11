//go:build !release

package util

import "os"

func init() {
	isDebug = true
	exeDirectory, _ = os.Getwd()
}
