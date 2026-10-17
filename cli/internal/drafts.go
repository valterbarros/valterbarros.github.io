package internal

import (
	"fmt"
	"valterbarros/blog-tools/internal/utils"
)

var draftsFolder = "../../articles/drafts"

func ListDrafts() {
	files, _ := utils.ListFiles(draftsFolder)

	fmt.Println("drafts: ", files)
}
