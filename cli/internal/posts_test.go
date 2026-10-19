package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"
	"valterbarros/blog-tools/internal/utils"
)

func TestCreatePosts(t *testing.T) {
	CreateNewPost()

	path, _ := utils.GetFromRoot(postsFolder)
	_, err := os.ReadFile(path + "/new-post.md")

	if err != nil {
		t.Error("The file was not created")
	}
}

func TestConvertToDraft(t *testing.T) {
	name := "post-test.md"
	err := ConvertToDraft(name)

	if err != nil {
		fmt.Println("an error occurred during convert", err)
	}
}

func TestFormatPost(t *testing.T) {
	t.Run("should add metadata", func(t *testing.T) {
		_, dir := createFile(t)

		FormatPost("post-test.md", dir)

		newPath := getRenamed(t, dir)
		data := utils.GetFileData(newPath)

		reg := regexp.MustCompile(`(?m)^title:|layout: post`)

		if len(reg.FindAllStringSubmatch(data, 2)) < 2 {
			t.Error("Not found metadata")
		}
	})

	t.Run("should add justify tags", func(t *testing.T) {
		_, dir := createFile(t)

		FormatPost("post-test.md", dir)

		newPath := getRenamed(t, dir)
		data := utils.GetFileData(newPath)

		fmt.Println("message", strings.Count(data, `<p align="justify">`))

		if strings.Count(data, `<p align="justify">`) > 3 {
			t.Error("Not found justify tags")
		}
	})
}

func getRenamed(t *testing.T, dir string) string {
	t.Helper()
	files, _ := utils.ListFiles(dir)
	index := slices.IndexFunc(files, func(name string) bool {
		return strings.Contains(name, "post-test.md")
	})
	newPath := filepath.Join(dir, files[index])
	return newPath
}

func createFile(t *testing.T) (string, string) {
	t.Helper()

	template := `---
layout: post
title: title
---

lorem ipsum testando .asodkasdaskdas 1trestasdasd
lorem ipsum testando .asodkasdaskdas 2trestasdasd
lorem ipsum testando .asodkasdaskdas 3trestasdasd
lorem ipsum testando .asodkasdaskdas 4trestasdasd
lorem ipsum testando .asodkasdaskdas 5trestasdasd
lorem ipsum testando .asodkasdaskdas 6trestasdasd.

lorem ipsum testando .asodkasdaskdas 7trestasdasd
lorem ipsum testando .asodkasdaskdas 8trestasdasd
lorem ipsum testando .asodkasdaskdas 9trestasdasd
lorem ipsum testando .asodkasdaskdas 10trestasdasd.

lorem ipsum testando .asodkasdaskdas 11trestasdasd
lorem ipsum testando .asodkasdaskdas 12trestasdasd
lorem ipsum testando .asodkasdaskdas 13trestasdasd.
`

	dir := t.TempDir()
	path := filepath.Join(dir, "post-test.md")
	utils.SaveFile(path, template)

	return path, dir
}
