package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ListFiles(path string) ([]string, error) {
	files, err := os.ReadDir(path)

	if err != nil {
		return []string{}, err
	}

	names := []string{}

	for _, f := range files {
		names = append(names, f.Name())
	}

	return names, nil
}

const repoName = "valterbarros.github.io"

func GetFromRoot(path string) (string, error) {
	pwd, _ := os.Getwd()

	if !strings.Contains(pwd, repoName) {
		return "", fmt.Errorf("go to right path")
	}

	base, _, _ := strings.Cut(pwd, repoName)

	return filepath.Join(base, repoName, path), nil
}

func SaveFile(file, data string) error {
	return os.WriteFile(file, []byte(data), 0644)
}

func ResetFile(file string) {
	os.WriteFile(file, []byte(""), 0644)
}

func GetFileData(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Error to open file utils.go", err)
	}

	return string(data)
}
