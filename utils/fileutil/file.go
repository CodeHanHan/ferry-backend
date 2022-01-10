package fileutil

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
)

const (
	JPEG = "jpeg"
	JPG  = "jpg"
	PNG  = "png"
	GIF  = "gif"
	TIF  = "tif"
	BMP  = "bmp"
	XLS  = "xls"
	XLSX = "xlsx"
	GO   = "golang"
)

var fileTypeMap sync.Map

func init() {
	fileTypeMap.Store("ffd8", JPEG)     // JPEG jpg
	fileTypeMap.Store("8950", PNG)      // PNG (png)
	fileTypeMap.Store("4749", GIF)      // GIF (gif)
	fileTypeMap.Store("4949", TIF)      // TIFF (tif)
	fileTypeMap.Store("4d4d", TIF)      // TIFF (tif)
	fileTypeMap.Store("424d", BMP)      // bmp
	fileTypeMap.Store("d0cf11e0", XLS)  // 2003 xls
	fileTypeMap.Store("504b0304", XLSX) // 2007 xlsx
	fileTypeMap.Store("7061", GO)       // 2007 xlsx
}

// 获取前二字节的十六进制
func bytesToHexString(src []byte) string {
	res := bytes.Buffer{}
	if src == nil || len(src) <= 0 {
		return ""
	}

	temp := make([]byte, 0)
	for i := 0; i < 2; i++ {
		v := src[i]
		sub := v & 0xFF
		hv := hex.EncodeToString(append(temp, sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0), 10))
		}
		res.WriteString(hv)
	}

	return res.String()
}

// GetFileType get the file type by the top several bytes
func GetFileType(f *os.File) (fileType string, err error) {
	data := make([]byte, 10)
	_, err = f.Read(data)
	if err != nil {
		return "", err
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		return "", err
	}

	fileCode := bytesToHexString(data)
	fileTypeMap.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(string)
		if strings.HasPrefix(fileCode, strings.ToLower(k)) ||
			strings.HasPrefix(k, strings.ToLower(fileCode)) {
			fileType = v
			return false
		}
		return true
	})

	return fileType, nil
}

// 判断文件的类型
// data: 文件流前十字节, 未找到类型fileType == ""
func GetFileTypeByFileBytes(data []byte) (fileType string) {
	fileCode := bytesToHexString(data)

	fileTypeMap.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(string)
		if strings.HasPrefix(fileCode, strings.ToLower(k)) ||
			strings.HasPrefix(k, strings.ToLower(fileCode)) {
			fileType = v
			return false
		}
		return true
	})
	return
}

// Copy copy a file from src to dst
func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	dstDir := dst[:strings.LastIndex(dst, "/")]
	if _, err := os.Stat(dstDir); err != nil {
		if os.IsNotExist(err) {
			if err := MakeDirAll(dstDir, os.ModePerm); err != nil {
				return 0, err
			}
		} else {
			return 0, err
		}
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func Split(filepath string) (dir, file, ext string) {
	dir, name := path.Split(filepath)
	ext = path.Ext(name)
	file = name[0 : len(name)-len(ext)]

	return
}
