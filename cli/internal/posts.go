package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"valterbarros/blog-tools/internal/utils"
)

var postsFolder = "articles/_posts"

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

	return os.Rename(newPostFile, draftFile)
}
