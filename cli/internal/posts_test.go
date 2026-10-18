package internal

import (
	"fmt"
	"os"
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
