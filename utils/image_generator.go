package utils

import (
	"errors"
	"image"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

type ImageGenerator struct {
	Img image.Image
	Dir string
	Ext string
}

func (this *ImageGenerator) PreGenerate(imagePath string) error {
	// Open the image.
	var err error
	this.Img, err = imaging.Open(imagePath)
	if err != nil {
		return err
	}

	this.Ext = strings.ToLower(filepath.Ext(imagePath))
	// validate
	if this.Ext != ".png" && this.Ext != ".jpg" && this.Ext != "jpeg" {
		return errors.New("Must be either png or jpg file")
	}

	b := this.Img.Bounds()
	if b.Max.X < 1024 || b.Max.Y < 1024 {
		return errors.New("width and height must be greater than 1024")
	}
	if b.Max.X != b.Max.Y {
		return errors.New("width and height must be same")
	}

	if imagePath[0:1] != "." {
		this.Dir = filepath.Dir(imagePath)
	} else {
		this.Dir = "./"
	}

	return nil
}

func (this *ImageGenerator) Generate() error {

	err := this.resize(20)
	if err != nil {
		return err
	}
	err = this.resize(29)
	if err != nil {
		return err
	}
	err = this.resize(40)
	if err != nil {
		return err
	}
	err = this.resize(58)
	if err != nil {
		return err
	}
	err = this.resize(76)
	if err != nil {
		return err
	}
	err = this.resize(80)
	if err != nil {
		return err
	}
	err = this.resize(87)
	if err != nil {
		return err
	}
	err = this.resize(120)
	if err != nil {
		return err
	}
	err = this.resize(152)
	if err != nil {
		return err
	}
	err = this.resize(167)
	if err != nil {
		return err
	}
	err = this.resize(180)
	if err != nil {
		return err
	}
	err = this.resize(1024)
	if err != nil {
		return err
	}
	return nil
}

func (this *ImageGenerator) resize(size int) error {
	src := imaging.Resize(this.Img, size, size, imaging.Lanczos)
	err := imaging.Save(src, this.Dir+"/Icon-"+strconv.Itoa(size)+this.Ext)
	if err != nil {
		return err
	}
	return nil
}
