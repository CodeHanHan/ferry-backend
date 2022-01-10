package imageutil

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path"
	"path/filepath"

	"github.com/nfnt/resize"
)

func ReadFrom(filename string) (image.Image, string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	return image.Decode(file)
}

func SaveJPEG(img image.Image, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	return jpeg.Encode(file, img, &jpeg.Options{100})
}

func ReSize(srcImgName string, width, height uint) (string, error) {
	srcImg, err := os.Open(srcImgName)
	if err != nil {
		return "", err
	}
	defer srcImg.Close()

	img, err := jpeg.Decode(srcImg)
	if err != nil {
		return "", err
	}

	m := resize.Resize(width, height, img, resize.Lanczos3)

	dir, name := path.Split(srcImgName)
	ext := path.Ext(name)
	file := name[0 : len(name)-len(ext)]
	dstName := filepath.Join(dir, fmt.Sprintf("%s-%d-%d%s", file, width, height, ext))

	out, err := os.Create(dstName)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// write new image to file
	err = jpeg.Encode(out, m, nil)
	if err != nil {
		return "", err
	}

	return dstName, nil
}
