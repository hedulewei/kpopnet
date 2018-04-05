package kpopnet

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"

	"github.com/Kagami/go-dlib"
)

const (
	minDimension = 300
	maxDimension = 5000
)

var (
	faceRec *dlib.FaceRec
)

func StartFaceRec(dataDir string) (err error) {
	faceRec, err = dlib.NewFaceRec(getModelsDir(dataDir))
	if err != nil {
		return fmt.Errorf("Error initializing face recognizer: %v", err)
	}
	return
}

// TODO(Kagami): Search for already recognized idol.
func Recognize(imgData []byte) (idolId *string, err error) {
	r := bytes.NewReader(imgData)
	c, typ, err := image.DecodeConfig(r)
	if err != nil || typ != "jpeg" ||
		c.Width < minDimension ||
		c.Height < minDimension ||
		c.Width > maxDimension ||
		c.Height > maxDimension {
		err = errBadImage
		return
	}
	_, err = faceRec.RecognizeSingle(imgData)
	if err != nil {
		return
	}
	return
}