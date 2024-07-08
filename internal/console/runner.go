// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package console

import (
	"errors"
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
	"github.com/guanguans/gh-actions-watcher/internal/git"
	"github.com/guanguans/gh-actions-watcher/internal/github/client"
	"github.com/guanguans/gh-actions-watcher/internal/github/entity"
)

const requestInterval = 5

type Runner struct {
	output *Output
	github *client.Github
	repo   string
	branch string
}

func NewDefaultRunner(repo, branch string) (*Runner, error) {
	gh, err := client.NewDefaultGithub()
	if err != nil {
		return nil, fmt.Errorf("failed to create Github client: %w", err)
	}

	localGitRepo, err := git.NewDefaultLocalGitRepo()
	if err != nil {
		return nil, fmt.Errorf("failed to create local git repo: %w", err)
	}

	if repo == "" {
		repo, err = localGitRepo.GetVendorAndRepo()
		if err != nil {
			return nil, fmt.Errorf("failed to get repo: %w", err)
		}
	}

	if branch == "" {
		branch, err = localGitRepo.GetCurrentBranch()
		if err != nil {
			return nil, fmt.Errorf("failed to get current branch: %w", err)
		}
	}

	return NewRunner(NewOutput(), gh, repo, branch), nil
}

func NewRunner(output *Output, github *client.Github, repo, branch string) *Runner {
	return &Runner{output: output, github: github, repo: repo, branch: branch}
}

func (r *Runner) Run() error {
	r.showHeader()
	r.output.LineInfo("Fetching Github workflow runs...")

	var lastWorkflows entity.WorkflowRunCollection

	for {
		workflows, err := r.displayWorkflows()
		if err != nil {
			return err
		}

		lastWorkflows = workflows

		if !r.shouldContinueWatching(workflows) {
			break
		}

	}

	if !lastWorkflows.AllCompletedSuccessfully() {
		return errors.New("some workflows failed")
	}

	r.output.BlockSuccess("All workflows finished successfully.")

	return nil
}

func (r *Runner) showHeader() {
	r.clearScreen()
	r.output.BlockInfo("Github Actions Watcher by guanguans")
}

func (r *Runner) clearScreen() {
	fmt.Print("\033\143")
}

func (r *Runner) showWorkflowRunCollection(runs entity.WorkflowRunCollection) {
	l := list.New().Enumerator(func(_ list.Items, _ int) string {
		return ""
	})

	for _, run := range runs.All() {
		status, _ := run.GetListStatus()

		render := lipgloss.NewStyle().
			Foreground(lipgloss.Color(status.Color())).
			Render(status.HumanReadableValue())

		l.Item(render + " ... " + run.Name)
	}

	r.output.NewLine(1)
	fmt.Println(l)
}

func (r *Runner) displayWorkflows() (entity.WorkflowRunCollection, error) {
	runs, err := r.github.LatestWorkflowRuns(r.repo, r.branch)
	if err != nil {
		return runs, fmt.Errorf("failed to get latest workflow runs: %w", err)
	}

	r.showHeader()
	r.output.LineInfo(fmt.Sprintf("Workflow runs for `%s` on the `%s` branch.", r.repo, r.branch))

	if runs.IsEmpty() {
		r.output.LineWarning("No workflow runs found for this repo...")

		return runs, nil
	}

	r.showWorkflowRunCollection(runs)

	if runs.ContainsActiveRuns() {
		r.output.BlockWarning("Running workflows detected. Refreshing automatically...")
	}

	return runs, nil
}

func (r *Runner) shouldContinueWatching(workflowRunCollection entity.WorkflowRunCollection) bool {
	if !workflowRunCollection.ContainsActiveRuns() {
		return false
	}

	time.Sleep(time.Second * requestInterval)

	return true
}
