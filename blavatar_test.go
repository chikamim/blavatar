package blavatar_test

import (
	"crypto/sha512"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"

	"github.com/chikamim/blavatar"
)

func ExampleNew() {
	img := blavatar.New("blavatar", 64)

	file := "blavatar.jpg"
	f, _ := os.Create(file)
	defer f.Close()
	jpeg.Encode(f, img, &jpeg.Options{70})
	got, _ := ioutil.ReadFile(file)

	fmt.Printf("%x", sha512.Sum384(got))
	// Output:
	// 6fc264f25e7110e39e5b354607a3b9d7eab80b5c07a8fa8babf82180b23b583ce83f603afd2a92bab3d785614cad3f68
}
