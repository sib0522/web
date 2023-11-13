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
	if os.IsNotExist(err) {
		// ディレクトリが存在しなかったら作成
		err = os.Mkdir(GalleryDirectory, os.ModeDir|os.ModePerm)
		if err != nil {
			return nil, err
		}
		// エラーをつぶす
		err = nil
	}
	if err != nil {
		return nil, err
	}

	list := make([]string, 0, len(files))

	for _, f := range files {
		//str := fmt.Sprintf("\".%v/%v\" class=\"w-100 p-lg-4 shadow-1-strong rounded mb-4\" alt=\"%v\"", constants.GalleryDirectory, f.Name(), f.Name())
		str := fmt.Sprintf(".%v/%v", GalleryDirectory, f.Name())
		list = append(list, str)
	}

	gallery := &Gallery{DataList: list}

	return gallery, nil
}
