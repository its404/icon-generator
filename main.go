package main

import (
	"fmt"
	"github/its404/icon-generator/utils"
)

func main() {
	var filePath string

	imageGenerator := utils.ImageGenerator{}
	for {
		fmt.Println("Please input file path(eg. /tmp/test.png):")
		fmt.Scanln(&filePath)
		err := imageGenerator.PreGenerate(filePath)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}
	if imageGenerator.Img != nil {
		err := imageGenerator.Generate()
		if err != nil {
			fmt.Printf("failed to generate image: %v", err)
		} else {
			fmt.Printf("images generated successfully")
		}
	} else {
		fmt.Println("Cannot get image data")
	}
}
