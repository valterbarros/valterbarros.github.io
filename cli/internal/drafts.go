package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"valterbarros/blog-tools/internal/utils"
)

var draftsFolder = "articles/drafts"

func ListDrafts() {
	path, _ := utils.GetFromRoot(draftsFolder)
	files, _ := utils.ListFiles(path)

	fmt.Println("drafts: ", files)
}

func CreateNewDraft() {
	path, _ := utils.GetFromRoot(draftsFolder)

	err := utils.SaveFile(filepath.Join(path, "new-draft.md"), "")

	if err != nil {
		fmt.Println("The file was not created: ", err)
	}
}

func ConvertToPost(name string) error {
	path, _ := utils.GetFromRoot(draftsFolder)
	postsPath, _ := utils.GetFromRoot(postsFolder)

	draftFile := filepath.Join(path, name)
	newPostFile := filepath.Join(postsPath, name)

	return os.Rename(draftFile, newPostFile)
}
