package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"valterbarros/blog-tools/internal/utils"
)

var draftsFolder = "articles/drafts"

func ListDrafts() (string, error) {
	path, err := utils.GetFromRoot(draftsFolder)
	files, err := utils.ListFiles(path)

	if err != nil {
		return "", err
	}

	return strings.Join(files, "\n"), nil
}

func CreateNewDraft(name string) error {
	if name == "" || filepath.Ext(name) == "" {
		name = "new-draft.md"
	}

	path, _ := utils.GetFromRoot(draftsFolder)

	err := utils.SaveFile(filepath.Join(path, name), "")

	if err != nil {
		fmt.Println("The file was not created: ", err)
		return err
	}

	return nil
}

func DeleteDraft(name string) error {
	path, _ := utils.GetFromRoot(draftsFolder)

	err := utils.RemoveFile(filepath.Join(path, name))

	if err != nil {
		fmt.Println("The file was not removed: ", err)
		return err
	}

	return nil
}

func ConvertToPost(name string) error {
	path, _ := utils.GetFromRoot(draftsFolder)
	postsPath, _ := utils.GetFromRoot(postsFolder)

	draftFile := filepath.Join(path, name)
	newPostFile := filepath.Join(postsPath, name)

	return os.Rename(draftFile, newPostFile)
}
