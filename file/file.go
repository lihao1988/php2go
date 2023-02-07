package file

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// Basename get file basename
// php basename
func Basename(fPath string) string {
	return filepath.Base(fPath)
}

// Dirname get file dirname
// php dirname
func Dirname(fPath string) string {
	return filepath.Dir(fPath)
}

// Filesize get file size
// php filesize
func Filesize(fPath string) int64 {
	fInfo, err := os.Stat(fPath)
	if err != nil && os.IsNotExist(err) {
		return 0
	}

	fm := fInfo.Mode()
	if fm.IsDir() {
		return 0
	}

	return fInfo.Size()
}

// PathInfo get file all info
// php pathinfo
func PathInfo(fPath string) (map[string]interface{}, error) {
	fInfoMap := map[string]interface{}{}
	fInfo, err := os.Stat(fPath)
	if err != nil && os.IsNotExist(err) {
		return fInfoMap, err
	}

	fm := fInfo.Mode()
	if fm.IsDir() {
		return fInfoMap, errors.New("fPath is dirname")
	}

	fInfoMap["dirname"] = filepath.Dir(fPath)
	fInfoMap["filename"] = fInfo.Name()
	fInfoMap["size"] = fInfo.Size()
	fInfoMap["extension"] = path.Ext(fPath)
	// fInfoMap["mode"] = fInfo.Mode()
	return fInfoMap, nil
}

// FileExists check file is exists
// php file_exists
func FileExists(fName string) bool {
	_, err := os.Stat(fName)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

// IsDir fileName is dir
// php is_dir
func IsDir(fName string) bool {
	fInfo, err := os.Stat(fName)
	if err != nil {
		return false
	}

	fm := fInfo.Mode()
	return fm.IsDir()
}

// IsFile fileName is file
// php is_file
func IsFile(fName string) bool {
	fInfo, err := os.Stat(fName)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	fm := fInfo.Mode()
	return !fm.IsDir()
}

// FileGetContents read file content
//php file_get_contents
func FileGetContents(fName string) (string, error) {
	data, err := ioutil.ReadFile(fName)
	return string(data), err
}

// FilePutContents write file content
// php file_put_contents
func FilePutContents(filename string, data string, mode os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), mode)
}

// Chmod chmod_file mode
// php chmod
func Chmod(filename string, mode os.FileMode) bool {
	return os.Chmod(filename, mode) == nil
}

// Chown chown_file
// php chown
func Chown(filename string, uid, gid int) bool {
	return os.Chown(filename, uid, gid) == nil
}
