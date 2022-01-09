package fileutil

import (
	"encoding/hex"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_bytesToHexString(t *testing.T) {
	f, err := os.Open("./path.go")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	data := make([]byte, 100)
	_, _ = f.Read(data)

	str := bytesToHexString(data)

	require.Equal(t, str, "7061")
}

func Test_GetFileType(t *testing.T) {
	f, err := os.Open("./path_test.go")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	fileType, err := GetFileType(f)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, GO, fileType)
}

func TestGetFileTypeByFileBytes(t *testing.T) {
	dataXls, _ := hex.DecodeString("d0cf11e0")
	dataXlsx, _ := hex.DecodeString("504b0304")

	require.Equal(t, "xls", GetFileTypeByFileBytes(dataXls))
	require.Equal(t, "xlsx", GetFileTypeByFileBytes(dataXlsx))
}
