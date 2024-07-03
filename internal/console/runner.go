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

type runner struct {
	github github.Github
	repo   string
	branch string
}

func NewDefaultRunner(repo string, branch string) (*runner, error) {
	defaultGithub, err := github.NewDefaultGithub()
	if err != nil {
		return nil, err
	}

	defaultLocalGitRepo, err := git.NewDefaultLocalGitRepo()
	if err != nil {
		return nil, err
	}

	if repo == "" {
		repo, err = defaultLocalGitRepo.GetVendorAndRepo()
		if err != nil {
			return nil, err
		}
	}

	if branch == "" {
		branch, err = defaultLocalGitRepo.GetCurrentBranch()
		if err != nil {
			return nil, err
		}
	}

	return newRunner(defaultGithub, repo, branch), nil
}

func newRunner(github *github.Github, repo string, branch string) *runner {
	return &runner{github: *github, repo: repo, branch: branch}
}

func (r runner) Hanlde() error {
	r.showHeader()

	NewConsoleOutput().lineInfo("Fetching GitHub workflow runs...")

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

	NewConsoleOutput().success("All workflows finished successfully.")

	return nil
}

func (r runner) displayWorkflows() (entity.WorkflowRunCollection, error) {
	runs, err := r.github.GetLatestWorkflowRuns(r.repo, r.branch)
	if err != nil {
		return runs, err
	}

	r.showHeader()

	NewConsoleOutput().lineInfo(fmt.Sprintf("Workflow runs for %s on the %s branch.\n", r.repo, r.branch))

	if runs.IsEmpty() {
		NewConsoleOutput().lineWarning("No workflow runs found for this repo...")
		return runs, nil
	}

	r.showRuns(runs)

	if runs.ContainsActiveRuns() {
		NewConsoleOutput().warning("Running workflows detected. Refreshing automatically...")
	}

	return runs, err
}

func (r runner) showRuns(runs entity.WorkflowRunCollection) {
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

func (r runner) showHeader() {
	r.clearScreen()

	NewConsoleOutput().info("GitHub Actions Watcher by guanguans - Logged in as guanguans")
}

func (r runner) clearScreen() {
	fmt.Print("\033\143")
}

func (r runner) shouldContinueWatching(workflowRunCollection entity.WorkflowRunCollection) bool {
	if !workflowRunCollection.ContainsActiveRuns() {
		return false
	}

	time.Sleep(8 * time.Second)

	return true
}
