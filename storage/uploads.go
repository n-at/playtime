package storage

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"regexp"
	"strings"
)

func (s *Storage) prepareUploadPath(id string) (string, error) {
	uploadPath, err := GetUploadPath(id)
	if err != nil {
		return "", err
	}

	uploadPath = fmt.Sprintf("%s%c%s", s.config.UploadsPath, os.PathSeparator, uploadPath)

	if err := os.MkdirAll(uploadPath, 0777); err != nil {
		return "", err
	}

	return uploadPath, nil
}

func (s *Storage) SaveUploadedFile(file *multipart.FileHeader, id, extension string) error {
	uploadPath, err := s.prepareUploadPath(id)
	if err != nil {
		return err
	}

	fileName := id
	if len(extension) != 0 {
		fileName += "." + extension
	}
	fileName = path.Join(uploadPath, fileName)

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func GetUploadPath(id string) (string, error) {
	if len(id) == 0 {
		return "", errors.New("id is empty")
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}

	filtered := reg.ReplaceAllString(id, "")

	var parts []string
	for i := 0; i < len(filtered); i += 2 {
		if i+2 <= len(filtered) {
			parts = append(parts, string([]rune(filtered)[i:i+2]))
		} else {
			parts = append(parts, string([]rune(filtered)[i:i+1]))
		}
	}

	return strings.Join(parts, string(os.PathSeparator)), nil
}
