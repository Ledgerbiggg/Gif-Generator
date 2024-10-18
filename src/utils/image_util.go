package utils

import (
	"errors"
	"gif_generator/src/common/consts"
	"github.com/disintegration/imaging"
	"golang.org/x/image/webp"
	"image"
	"os"
	"path/filepath"
)

// ImageDecode to decode images
func ImageDecode(file string) (image.Image, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	ext := filepath.Ext(file)
	switch ext {
	case consts.BMP, consts.GIF, consts.JPG, consts.JPEG, consts.PNG:
		return imaging.Decode(f)
	case consts.WEBP:
		return webp.Decode(f)
	default:
		return nil, errors.New("unsupported images format")
	}

}
