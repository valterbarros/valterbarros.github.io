package internal

import (
	"fmt"
	"path/filepath"
	"strings"
	"valterbarros/blog-tools/internal/utils"
)

var postsFolder = "articles/_posts"

// TODO: try generic methods with drafts code

func ListPosts() (string, error) {
	path, err := utils.GetFromRoot(postsFolder)
	files, err := utils.ListFiles(path)

	if err != nil {
		return "", err
	}

	return strings.Join(files, "\n"), nil
}

func CreateNewPost() error {
	path, _ := utils.GetFromRoot(postsFolder)

	err := utils.SaveFile(filepath.Join(path, "new-post.md"), "")

	if err != nil {
		fmt.Println("The file was not created: ", err)

		return err
	}

	return nil
}

func DeletePost(name string) error {
	path, _ := utils.GetFromRoot(postsFolder)

	err := utils.RemoveFile(filepath.Join(path, name))

	if err != nil {
		fmt.Println("The file was not removed: ", err)
		return err
	}

	return nil
}

func ConvertToDraft(name string) error {
	path, _ := utils.GetFromRoot(draftsFolder)
	postsPath, _ := utils.GetFromRoot(postsFolder)

	newPostFile := filepath.Join(postsPath, name)
	draftFile := filepath.Join(path, name)

	return utils.RenameFile(newPostFile, draftFile)
}

// specific methods
func FormatPost(name string, folder string) error {
	if folder == "" {
		folder = postsFolder
	}

	now := utils.GetDate()

	postsPath, _ := utils.GetFromRoot(folder)

	currentPostFile := filepath.Join(postsPath, name)
	renamedPostFile := filepath.Join(postsPath, now+"-"+name)

	metadata := `---
layout: post
title: [title]
---`

	data := utils.GetFileData(currentPostFile)

	justified := addJustify(data)
	fmt.Println(justified)

	utils.SaveFile(currentPostFile, metadata+"\n"+justified)

	return utils.RenameFile(currentPostFile, renamedPostFile)
}

// isOpen logic dont work when there is two paragraph followed
// isPara logic dont work when paragraph is not followed like has a title in the middle
// startPara dont work the start index get outdated

func addJustify(data string) string {
	strs := strings.Split(strings.TrimSpace(data), "\n")
	formmated := ""
	left := 0

	for left < len(strs) {
		line := strs[left]
		if isEmpty(line) {
			formmated += mountPara(strs, left)
			left += 1
		} else {
			formmated += line + "\n"
		}

		left += 1
	}

	return formmated
}

func mountPara(strs []string, left int) string {
	if !validateLine(strs, left) {
		return "\n" + strs[left+1] + "\n"
	}

	return `
<p align="justify">
	` + strs[left+1] + "\n</p>\n"
}

// TODO: create tests for that
func validateLine(strs []string, idx int) bool {
	// first element is a empty character idx + 1
	newI := idx + 1

	return len(strings.TrimSpace(strs[newI])) > 100
}

func isEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}
