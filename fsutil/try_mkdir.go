package fsutil

import "os"

func TryCreatePath(filePath string) error {
	return os.MkdirAll(filePath, 0750)
}
