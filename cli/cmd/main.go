package main

import (
	"fmt"
	"log"
	"strings"
	"valterbarros/blog-tools/internal/content"
	"valterbarros/blog-tools/internal/editor"
	"valterbarros/blog-tools/internal/utils"

	"github.com/spf13/cobra"
)

func draftCommands() *cobra.Command {
	var listDraft bool
	var createDraft bool
	var convertDraft bool
	var removeDraft bool

	listDraftCmd := &cobra.Command{
		Use:   "drafts",
		Short: "utils to create/list/delete posts.",
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if listDraft {
				d, err := content.ListDrafts()
				fmt.Println(d)

				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
			}

			if createDraft {
				err := content.CreateNewDraft("")

				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
			}

			if convertDraft {
				err := content.ConvertToPost(args[0])

				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
			}

			if removeDraft {
				err := content.DeleteDraft(args[0])

				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
			}

			return nil
		},
	}

	listDraftCmd.Flags().BoolVarP(&listDraft, "list", "l", false, "List all drafts files")
	listDraftCmd.Flags().BoolVarP(&createDraft, "create", "c", false, "Create a draft file")
	listDraftCmd.Flags().BoolVarP(&removeDraft, "delete", "d", false, "Remove a draft file")
	listDraftCmd.Flags().BoolVarP(&convertDraft, "ready", "r", false, "Convert a draft to a post file")

	return listDraftCmd
}

func postsCommands() *cobra.Command {
	var listPost bool
	var createPost bool
	var convertPost bool
	var removePost bool
	var formatPost bool
	var editPost bool

	postCmd := &cobra.Command{
		Use:   "posts",
		Short: "utils to create/list/delete posts.",
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if listPost {
				d, err := content.ListPosts()
				fmt.Println(d)

				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
			}

			if createPost {
				err := content.CreateNewPost("")

				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
			}

			if convertPost {
				err := content.ConvertToDraft(args[0])

				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
			}

			if removePost {
				err := content.DeletePost(args[0])

				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
			}

			if formatPost {
				err := content.FormatPost(args[0], "")

				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
			}

			if editPost {
				err := editor.RunServer()

				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
			}

			return nil
		},
	}

	postCmd.Flags().BoolVarP(&listPost, "list", "l", false, "List all posts files")
	postCmd.Flags().BoolVarP(&createPost, "create", "c", false, "Create a post file")
	postCmd.Flags().BoolVarP(&removePost, "delete", "r", false, "Remove a post file")
	postCmd.Flags().BoolVarP(&convertPost, "draft", "d", false, "Convert a post to a draft file")
	postCmd.Flags().BoolVarP(&formatPost, "format", "f", false, "Format post file to be published")
	postCmd.Flags().BoolVarP(&editPost, "editor", "e", false, "Edit a post file")

	return postCmd
}

func commitCommands() *cobra.Command {
	return &cobra.Command{
		Use:   "commit",
		Short: "commit data.",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			changes := utils.GetChanges()

			if strings.Contains(changes, "articles") {
				path, _ := utils.GetFromRoot("articles")
				utils.CommitChanges("[cli blogging] articles updated", path)
			} else {
				fmt.Println("not staged articles/**/* folder")
			}
		},
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "blogging",
		Short: "Manage blog CLI.",
	}

	dC := draftCommands()
	pC := postsCommands()
	cC := commitCommands()

	rootCmd.AddCommand(dC)
	rootCmd.AddCommand(pC)
	rootCmd.AddCommand(cC)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
