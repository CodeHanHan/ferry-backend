package imageutil

import (
	"fmt"
	"os"

	"github.com/h2non/bimg"
)

type Size [2]int

var (
	Size60 Size = [2]int{60, 60}
	// ...
)

func ReSize(srcImg []byte, dstSize Size) (dstImg []byte, err error) {
	newImage, err := bimg.NewImage(srcImg).Resize(dstSize[0], dstSize[1])
	if err != nil {
		return nil, err
	}

	return newImage, nil
}

func ForceReSize(srcImg []byte, dstSize Size) (dstImg []byte, err error) {
	newImage, err := bimg.NewImage(srcImg).ForceResize(dstSize[0], dstSize[1])
	if err != nil {
		return nil, err
	}

	return newImage, nil
}

func Convert2Png(srcImg []byte) (dstImg []byte, err error) {
	newImage, err := bimg.NewImage(srcImg).Convert(bimg.PNG)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	return newImage, nil
}
