package main

import (
	"fmt"
	"log"
	"strings"

	"valterbarros/blog-tools/internal"
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
		Short: "drafts.",
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if listDraft {
				d, err := internal.ListDrafts()
				fmt.Println(d)

				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
			}

			if createDraft {
				err := internal.CreateNewDraft("")

				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
			}

			if convertDraft {
				err := internal.ConvertToPost(args[0])

				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
			}

			if removeDraft {
				err := internal.DeleteDraft(args[0])

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

	postCmd := &cobra.Command{
		Use:   "posts",
		Short: "posts.",
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if listPost {
				d, err := internal.ListPosts()
				fmt.Println(d)

				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
			}

			if createPost {
				err := internal.CreateNewPost()

				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
			}

			if convertPost {
				err := internal.ConvertToDraft(args[0])

				if err != nil {
					return fmt.Errorf("error: %w", err)
				}
			}

			if removePost {
				err := internal.DeletePost(args[0])

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
