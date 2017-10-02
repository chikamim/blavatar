package blavatar

import (
	"crypto/sha512"
	"image"
	"image/color"
	"math"

	"github.com/disintegration/imaging"
)

// New returns blavatar image from string
func New(str string, size int) *image.NRGBA {
	hash := sha512.Sum384([]byte(str))
	pixels := bytesToPixels(hash)
	image := pixelsToImage(pixels)
	image = blurImage(image, size)
	return image
}

func bytesToPixels(bin [48]byte) (pixels []color.NRGBA) {
	for p := 0; p < 16; p++ {
		rgb := bin[p*3 : p*3+3]
		pixel := color.NRGBA{rgb[0], rgb[1], rgb[2], 255}
		pixels = append(pixels, pixel)
	}
	return pixels
}

func pixelsToImage(pixels []color.NRGBA) *image.NRGBA {
	size := int(math.Sqrt(float64(len(pixels))))
	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	for i, p := range pixels {
		x := i % size
		y := i / size
		img.Set(x, y, p)
	}
	return img
}

func blurImage(img *image.NRGBA, size int) *image.NRGBA {
	img = imaging.Resize(img, size, size, imaging.Gaussian)
	img = imaging.AdjustBrightness(img, (float64)(size/4))
	img = imaging.Blur(img, (float64)(size/8))
	return img
}
