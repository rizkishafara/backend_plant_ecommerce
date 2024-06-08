package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetFileTypeFromBase64Header(header string) (string, error) {
	switch {
	case strings.Contains(header, "image/jpeg"):
		return "jpeg", nil
	case strings.Contains(header, "image/jpg"):
		return "jpg", nil
	case strings.Contains(header, "image/png"):
		return "png", nil
	default:
		return "", errors.New("unsupported file type")
	}
}
func SaveFile(newFileName, fileType, jenis string, file []byte) (string, error) {

	path := fmt.Sprintf("/uploads/%s", jenis)
	filePath := fmt.Sprintf("/%s", newFileName)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			return "", errors.New("unable to create folder")
		}
	}
	if err := ioutil.WriteFile(path+filePath, file, 0644); err != nil {
		return "", errors.New("unable to save file")
	}
	return filePath, nil

}

func DeleteFile(file, jenis string) error {
	path := fmt.Sprintf("%s/%s", jenis, file)
	err := os.Remove("/uploads/" + path)
	if err != nil {
		return err
	}
	return nil
}
