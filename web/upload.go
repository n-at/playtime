package web

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func (s *Server) prepareUploadPath(id string) (string, error) {
	uploadPath, err := getUploadPath(id)
	if err != nil {
		return "", err
	}

	uploadPath = fmt.Sprintf("%s%c%s", s.config.UploadsRoot, os.PathSeparator, uploadPath)

	if err := os.MkdirAll(uploadPath, 0777); err != nil {
		return "", err
	}

	return uploadPath, nil
}

func getUploadPath(id string) (string, error) {
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

func getFileExtension(name string) string {
	if len(name) == 0 {
		return ""
	}

	parts := strings.Split(name, ".")
	if len(parts) == 1 {
		return ""
	}

	return parts[len(parts)-1]
}
