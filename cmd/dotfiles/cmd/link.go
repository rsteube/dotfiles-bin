package cmd

import (
	"github.com/mitchellh/go-homedir"
	"github.com/rsteube/carapace"
	"github.com/rsteube/dotfiles-bin/pkg/dotfiles"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "link a file",
	Run: func(cmd *cobra.Command, args []string) {
		dotfiles.Walk(func(d dotfiles.Dotfile) error {
			if err := d.Symlink(false); err != nil {
				println(err.Error())
			}
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)

	carapace.Gen(editCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(args []string) carapace.Action {
			if dotfiledir, err := homedir.Expand("~/.local/share/dotfiles"); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				return ActionSubDirectoryFiles(dotfiledir)
			}
		}),
	)
}
