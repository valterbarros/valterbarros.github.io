package utils

import (
	"strings"
	"testing"
)

var draftsFolder = "../../articles/_posts"

func TestFiles(t *testing.T) {
	path, _ := GetFromRoot("articles/_posts")
	files, _ := ListFiles(path)

	if len(files) < 1 {
		t.Error("The directory is empty")
	}
}

func TestGoFromRoot(t *testing.T) {
	path, _ := GetFromRoot("new_path/other")

	if !strings.Contains(path, "new_path/other") {
		t.Error("The path is wrong")
	}
}
