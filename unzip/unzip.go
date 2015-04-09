package unzip

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path"
)

const (
	TMP_DIRECTORY = "tmp/"
)

func Unzip(src string) error {
	if err := TmpDirectoryCreate(); err != nil {
		log.Println(err)
		return err
	}

	r, err := zip.OpenReader(src)

	if err != nil {
		log.Println(err)
		return err
	}

	defer r.Close()

	for _, file := range r.File {
		rc, err := file.Open()

		if err != nil {
			log.Println(err)
			return err
		}

		defer rc.Close()

		file_path := path.Join(TMP_DIRECTORY, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(file_path, file.Mode())
		} else {
			newFile, err := os.Create(file_path)

			if err != nil {
				log.Println(err)
				return err
			}

			io.Copy(newFile, rc)
			newFile.Sync()
		}
	}
	return nil
}

func TmpDirectoryCreate() error {
	return os.MkdirAll(TMP_DIRECTORY, 0775)
}
