package file

import (
	"fmt"
	"testing"
)

func TestFileFunc(t *testing.T) {
	fPath := "/home/anyuan/下载/php2go-1.0.0/go.mod"
	fmt.Println("Basename: ", Basename(fPath))

	fmt.Println("Dirname: ", Dirname(fPath))

	fmt.Println("Filesize: ", Filesize(fPath))

	fInfo, err := PathInfo(fPath)
	fmt.Println("PathInfo: ", fInfo, err)

	fmt.Println("FileExists: ", FileExists(fPath))

	fPathDir := "/home/anyuan/下载/php2go-1.0.0/"
	fmt.Println("IsDir - IsFile: ", IsDir(fPathDir), IsFile(fPathDir))
	fmt.Println("IsFile - IsDir: ", IsFile(fPath), IsDir(fPath))

	fmt.Println(FileGetContents(fPath))

}
