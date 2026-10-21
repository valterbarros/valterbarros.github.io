package internal

import (
	"fmt"
	"path/filepath"
	"regexp"
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

	justified := addJustifyTag(data)
	fmt.Println(justified)

	utils.SaveFile(currentPostFile, metadata+"\n"+justified)

	return utils.RenameFile(currentPostFile, renamedPostFile)
}

func addJustifyTag(data string) string {
	strs := strings.Split(strings.TrimSpace(data), "\n")
	formmated := ""
	left := 0

	for left < len(strs) {
		line := strs[left]
		if isEmpty(line) {
			// increment because current left is a empty character
			left += 1
			formmated += mountPara(strs, left)
		} else {
			formmated += line + "\n"
		}

		left += 1
	}

	return formmated
}

func mountPara(strs []string, left int) string {
	if !validateLine(strs, left) {
		return "\n" + strs[left] + "\n"
	}

	withTags := addHtmlTags(strs[left])

	return `
<p align="justify">
` + withTags + "\n</p>\n"
}

// TODO: create tests for that
func validateLine(strs []string, idx int) bool {
	newI := idx
	nextLine := newI + 1

	// Is end of file or next line is a \n
	isNextLineValid := nextLine >= len(strs) || strs[nextLine] == ""

	return len(strings.TrimSpace(strs[newI])) > 100 && isNextLineValid
}

func isEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

func addHtmlTags(str string) string {
	regLink := regexp.MustCompile(`(?mi)\[([^\]]+)\]\(([^\)]+)\)`)
	regAsterisk := regexp.MustCompile(`(?mi)\*{1,2}([^\*]+)\*{1,2}`)
	repl := regAsterisk.ReplaceAllStringFunc(str, func(s string) string {
		if strings.Count(s, "*") == 4 {
			return "<b>" + strings.ReplaceAll(s, "*", "") + "</b>"
		}

		return "<i>" + strings.ReplaceAll(s, "*", "") + "</i>"
	})
	repl = regLink.ReplaceAllString(repl, "<a href=\"$2\">$1</a>")

	return repl
}
