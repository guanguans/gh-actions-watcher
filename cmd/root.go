// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package cmd

import (
	"fmt"

	"github.com/guanguans/gh-actions-watcher/internal/console"
	"github.com/spf13/cobra"
)

func Execute() {
	var (
		repo    string
		branch  string
		rootCmd = &cobra.Command{
			Use:           "actions-watcher",
			Short:         "Watch the Github actions of a repo.",
			SilenceErrors: true,
			SilenceUsage:  true,
			RunE: func(_ *cobra.Command, _ []string) error {
				runner, err := console.NewDefaultRunner(repo, branch)
				if err != nil {
					return fmt.Errorf("failed to create runner: %w", err)
				}

				return runner.Run()
			},
		}
	)

	rootCmd.Flags().StringVarP(&repo, "repo", "r", "", "Github repository.")
	rootCmd.Flags().StringVarP(&branch, "branch", "b", "", "Workflow run branch.")

	if err := rootCmd.Execute(); err != nil {
		console.NewOutput().BlockError(err.Error())
	}
}
