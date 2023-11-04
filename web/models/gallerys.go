package models

import (
	"fmt"
	"os"
)

type Gallery struct {
	DataList []string
}

const GalleryDirectory = "./web/public/images"

func CreateGallery() (*Gallery, error) {
	files, err := os.ReadDir(GalleryDirectory)
	if err != nil {
		return nil, err
	}

	list := make([]string, 10)

	for i, f := range files {
		//str := fmt.Sprintf("\".%v/%v\" class=\"w-100 p-lg-4 shadow-1-strong rounded mb-4\" alt=\"%v\"", constants.GalleryDirectory, f.Name(), f.Name())
		str := fmt.Sprintf(".%v/%v", GalleryDirectory, f.Name())
		list[i] = str
	}

	gallery := &Gallery{DataList: list}

	return gallery, nil
}
