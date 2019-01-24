package utils

import (
	"errors"
	"image"
	"path/filepath"
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

	src := imaging.Resize(this.Img, 1024, 1024, imaging.Lanczos)
	err := imaging.Save(src, this.Dir+"/icon-1024"+this.Ext)
	if err != nil {
		return err
	}
	src = imaging.Resize(this.Img, 120, 120, imaging.Lanczos)
	err = imaging.Save(src, this.Dir+"/Icon-60@2x"+this.Ext)
	err = imaging.Save(src, this.Dir+"/Icon-Small-40@3x"+this.Ext)
	if err != nil {
		return err
	}
	src = imaging.Resize(this.Img, 180, 180, imaging.Lanczos)
	err = imaging.Save(src, this.Dir+"/Icon-60@3x"+this.Ext)
	if err != nil {
		return err
	}
	src = imaging.Resize(this.Img, 76, 76, imaging.Lanczos)
	err = imaging.Save(src, this.Dir+"/Icon-76"+this.Ext)
	if err != nil {
		return err
	}
	src = imaging.Resize(this.Img, 152, 152, imaging.Lanczos)
	err = imaging.Save(src, this.Dir+"/Icon-152"+this.Ext)
	if err != nil {
		return err
	}
	src = imaging.Resize(this.Img, 167, 167, imaging.Lanczos)
	err = imaging.Save(src, this.Dir+"/Icon-83.5@2x"+this.Ext)
	if err != nil {
		return err
	}
	src = imaging.Resize(this.Img, 40, 40, imaging.Lanczos)
	err = imaging.Save(src, this.Dir+"/Icon-Small-40"+this.Ext)
	if err != nil {
		return err
	}
	src = imaging.Resize(this.Img, 80, 80, imaging.Lanczos)
	err = imaging.Save(src, this.Dir+"/Icon-Small-40@2x"+this.Ext)
	if err != nil {
		return err
	}
	src = imaging.Resize(this.Img, 120, 120, imaging.Lanczos)
	err = imaging.Save(src, this.Dir+"/Icon-Small-40@3x"+this.Ext)
	if err != nil {
		return err
	}
	src = imaging.Resize(this.Img, 29, 29, imaging.Lanczos)
	err = imaging.Save(src, this.Dir+"/Icon-29"+this.Ext)
	if err != nil {
		return err
	}
	src = imaging.Resize(this.Img, 58, 58, imaging.Lanczos)
	err = imaging.Save(src, this.Dir+"/Icon-58"+this.Ext)
	if err != nil {
		return err
	}
	src = imaging.Resize(this.Img, 87, 87, imaging.Lanczos)
	err = imaging.Save(src, this.Dir+"/Icon-87"+this.Ext)
	if err != nil {
		return err
	}
	return nil
}
