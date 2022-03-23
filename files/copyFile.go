package files

import (
	"io"
	"mime/multipart"
	"os"
)

func CopyFile(targetFileName string, content multipart.File) (bool, error) {
	f, err := os.OpenFile(targetFileName, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return false, err
	}

	_, err = io.Copy(f, content)

	if err != nil {
		return false, err
	}
	return true, nil
}
