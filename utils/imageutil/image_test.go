package imageutil

// import (
// 	"io/ioutil"
// 	"os"
// 	"testing"

// 	"github.com/stretchr/testify/require"

// 	"github.com/CodeHanHan/ferry-backend/utils/fileutil"
// )

// const (
// 	srcImgName = "./1.jpeg"
// )

// func Test_ReSize(t *testing.T) {
// 	srcImg, err := ioutil.ReadFile(srcImgName)
// 	require.NoError(t, err)

// 	newImg, err := ReSize(srcImg, Size60)
// 	require.NoError(t, err)

// 	err = ioutil.WriteFile("2.jpeg", newImg, os.ModePerm)
// 	require.NoError(t, err)
// }

// func Test_ForceReSize(t *testing.T) {
// 	srcImg, err := ioutil.ReadFile(srcImgName)
// 	require.NoError(t, err)

// 	newImg, err := ReSize(srcImg, Size([2]int{1, 400}))
// 	require.NoError(t, err)

// 	err = ioutil.WriteFile("2.jpeg", newImg, os.ModePerm)
// 	require.NoError(t, err)
// }

// func Test_Convert2Png(t *testing.T) {
// 	srcImg, err := ioutil.ReadFile(srcImgName)
// 	require.NoError(t, err)

// 	newImg, err := Convert2Png(srcImg)
// 	require.NoError(t, err)

// 	getType := fileutil.GetFileTypeByFileBytes(newImg)
// 	require.Equal(t, getType, fileutil.PNG)

// 	err = ioutil.WriteFile("2.png", newImg, os.ModePerm)
// 	require.NoError(t, err)
// }
