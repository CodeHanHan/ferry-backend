package fileutil

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_file(t *testing.T) {
	hashDir := MakeHashDir("ferry")

	t.Log(hashDir)

	dir := fmt.Sprintf("./%s", hashDir)

	require.NoError(t, MakeDirAll(dir, os.ModePerm))

	dir_list := strings.Split(dir, "/")

	require.NoError(t, RemoveAll([]string{dir}))

	dir1 := strings.Join(dir_list[0:len(dir_list)-1], "/")

	require.NoError(t, RemoveAll([]string{dir1}))

	dir2 := strings.Join(dir_list[0:len(dir_list)-2], "/")

	require.NoError(t, RemoveAll([]string{dir2}))

}
