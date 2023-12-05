package finder

import "os"

func isSymlink(path string) (bool, string) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return false, ""
	}
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		linc, err := os.Readlink(path)
		if err != nil {
			return false, ""
		}
		return true, linc
	}
	return false, ""
}
