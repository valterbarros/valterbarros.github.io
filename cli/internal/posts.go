package internal

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"valterbarros/blog-tools/internal/utils"
)

var postsFolder = "articles/_posts"

var mn = NewManage(postsFolder)

func ListPosts() (string, error) {
	return mn.ListSome()
}

func CreateNewPost(name string) error {
	return mn.CreateSome(name)
}

func DeletePost(name string) error {
	return mn.DeleteSome(name)
}

func ConvertToDraft(name string) error {
	return mn.MoveSome(name, draftsFolder)
}

// specific methods
func FormatPost(name string, folder string) error {
	if folder == "" {
		folder = postsFolder
	}

	postsPath, _ := utils.GetFromRoot(folder)
	currentPostFile := filepath.Join(postsPath, name)
	data := utils.GetFileData(currentPostFile)

	justified := addJustifyTag(data)
	metadata := addMetadata(justified)
	err := utils.SaveFile(currentPostFile, metadata+"\n"+justified)
	if err != nil {
		fmt.Println("error: ", err)
	}

	newName := addDateToName(postsPath, name)

	return utils.RenameFile(currentPostFile, newName)
}

func addMetadata(justified string) string {
	if strings.Contains(justified, "layout:") {
		return ""
	}
	metadata := `---
layout: post
title: [title]
---`

	return metadata
}

func addDateToName(postsPath, name string) string {
	now := utils.GetDate()

	renamedPostFile := filepath.Join(postsPath, now+"-"+name)
	return renamedPostFile
}

func addJustifyTag(data string) string {
	strs := strings.Split(strings.TrimSpace(data), "\n")
	formmated := ""
	left := 0

	for left < len(strs) {
		line := strs[left]
		if isEmpty(line) {
			// increment because current left is an empty character
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

	// inside p tag the markdown don't work
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
