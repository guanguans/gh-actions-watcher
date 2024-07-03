// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package console

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
	"github.com/guanguans/gh-actions-watcher/internal/git"
	"github.com/guanguans/gh-actions-watcher/internal/github"
	"github.com/guanguans/gh-actions-watcher/internal/github/entity"
)

type Runner struct {
	output *Output
	github *github.Github
	repo   string
	branch string
}

func NewDefaultRunner(repo string, branch string) (*Runner, error) {
	gh, err := github.NewDefaultGithub()
	if err != nil {
		return nil, err
	}

	localGitRepo, err := git.NewDefaultLocalGitRepo()
	if err != nil {
		return nil, err
	}

	if repo == "" {
		repo, err = localGitRepo.GetVendorAndRepo()
		if err != nil {
			return nil, err
		}
	}

	if branch == "" {
		branch, err = localGitRepo.GetCurrentBranch()
		if err != nil {
			return nil, err
		}
	}

	return NewRunner(NewOutput(), gh, repo, branch), nil
}

func NewRunner(output *Output, github *github.Github, repo string, branch string) *Runner {
	return &Runner{output: output, github: github, repo: repo, branch: branch}
}

func (r *Runner) Run() error {
	r.showHeader()

	r.output.LineInfo("Fetching GitHub workflow runs...")

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
		return fmt.Errorf("Some workflows failed...")
	}

	r.output.Success("All workflows finished successfully.")

	return nil
}

func (r *Runner) showHeader() {
	r.clearScreen()
	r.output.Info("GitHub Actions Watcher by guanguans - Logged in as guanguans")
}

func (r *Runner) clearScreen() {
	fmt.Print("\033\143")
}

func (r *Runner) showWorkflowRunCollection(runs entity.WorkflowRunCollection) {
	l := list.New().Enumerator(func(items list.Items, index int) string {
		return ""
	})

	for _, run := range runs.All() {
		status, _ := run.GetListStatus()

		render := lipgloss.NewStyle().
			Foreground(lipgloss.Color(status.Color())).
			Render(status.HumanReadableValue())

		l.Item(render + " ... " + run.Name)
	}

	fmt.Println(l)
}

func (r *Runner) displayWorkflows() (entity.WorkflowRunCollection, error) {
	runs, err := r.github.GetLatestWorkflowRuns(r.repo, r.branch)
	if err != nil {
		return runs, err
	}

	r.showHeader()
	r.output.LineInfo(fmt.Sprintf("Workflow runs for %s on the %s branch.\n", r.repo, r.branch))

	if runs.IsEmpty() {
		r.output.LineWarning("No workflow runs found for this repo...")

		return runs, nil
	}

	r.showWorkflowRunCollection(runs)

	if runs.ContainsActiveRuns() {
		r.output.Warning("Running workflows detected. Refreshing automatically...")
	}

	return runs, err
}

func (r *Runner) shouldContinueWatching(workflowRunCollection entity.WorkflowRunCollection) bool {
	if !workflowRunCollection.ContainsActiveRuns() {
		return false
	}

	time.Sleep(time.Second * 5)

	return true
}
