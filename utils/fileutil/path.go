package fileutil

import (
	"fmt"
	"hash/crc32"
	"os"
	"path/filepath"
	"strings"
)

// 根据文件名生成目录
func MakeHashDir(key string) (hashDir string) {
	if key == "" {
		return ""
	}

	var hashNumber int
	hashNumber = int(crc32.ChecksumIEEE([]byte(key)))
	if -hashNumber >= 0 {
		hashNumber = -hashNumber
	}

	return fmt.Sprintf("%d/%d/%d", hashNumber&0xf, (hashNumber&0xf0)>>4, (hashNumber&0xf00)>>8)
}

// 生成多级文件目录 mkdir -p
func MakeDirAll(dir string, perm os.FileMode) error {
	_, err := os.Stat(dir) // 查看目录是否存在
	if err == nil {        // 若存在则直接退出
		return nil
	}

	if os.IsNotExist(err) {
		if err := os.MkdirAll(dir, perm); err != nil {
			return err
		}
		return nil
	}

	return err
}

func FindFilesByPrefix(dirPath, prefix string) (files []string, err error) {
	files = make([]string, 0, 5)

	err = filepath.Walk(dirPath, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		if strings.HasPrefix(fi.Name(), prefix) {
			files = append(files, filename)
		}
		return nil
	})

	return files, err
}

func RemoveAll(fileNames []string) error {
	for _, fileName := range fileNames {
		if err := os.Remove(fileName); err != nil {
			return err
		}
	}

	return nil
}
