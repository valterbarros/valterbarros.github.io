package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"valterbarros/blog-tools/internal/utils"
)

var postsFolder = "articles/_posts"

func ListPosts() {
	path, _ := utils.GetFromRoot(postsFolder)
	files, _ := utils.ListFiles(path)

	fmt.Println("posts: ", files)
}

func CreateNewPost() {
	path, _ := utils.GetFromRoot(postsFolder)

	err := utils.SaveFile(filepath.Join(path, "new-post.md"), "")

	if err != nil {
		fmt.Println("The file was not created: ", err)
	}
}

func ConvertToDraft(name string) error {
	path, _ := utils.GetFromRoot(draftsFolder)
	postsPath, _ := utils.GetFromRoot(postsFolder)

	newPostFile := filepath.Join(postsPath, name)
	draftFile := filepath.Join(path, name)

	return os.Rename(newPostFile, draftFile)
}
