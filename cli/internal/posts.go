package internal

import (
	"fmt"
	"valterbarros/blog-tools/internal/utils"
)

var postsFolder = "../../articles/_posts"

func ListPosts() {
	files, _ := utils.ListFiles(draftsFolder)

	fmt.Println("posts: ", files)
}
