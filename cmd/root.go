// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	var (
		repo    string
		branch  string
		rootCmd = &cobra.Command{
			Use:           "actions-watcher",
			Short:         "Watch the GitHub actions of a repo.",
			SilenceErrors: true,
			SilenceUsage:  true,
			RunE: func(cmd *cobra.Command, args []string) error {
				defaultRunner, err := newDefaultRunner(repo, branch)
				if err != nil {
					return err
				}

				return defaultRunner.hanlde()
			},
		}
	)

	rootCmd.Flags().StringVarP(&repo, "repo", "r", "", "GitHub repository.")
	rootCmd.Flags().StringVarP(&branch, "branch", "b", "", "Workflow run branch.")

	err := rootCmd.Execute()
	if err != nil {
		newConsoleOutput().error(err.Error())
		os.Exit(1)
	}
}
