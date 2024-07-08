// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package client

import (
	"fmt"
	"strings"

	gh "github.com/cli/go-gh/v2/pkg/api"
	"github.com/guanguans/gh-actions-watcher/internal/github/entity"
)

type Github struct {
	client *gh.RESTClient
}

func NewDefaultGithub() (*Github, error) {
	client, err := gh.DefaultRESTClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create GitHub client: %w", err)
	}

	return NewGithub(client), nil
}

func NewGithub(client *gh.RESTClient) *Github {
	return &Github{client: client}
}

func (g *Github) LatestWorkflowRuns(repo, branch string) (entity.WorkflowRunCollection, error) {
	workflowRunCollection, err := g.WorkflowRuns(repo, branch)
	if err != nil {
		return workflowRunCollection, err
	}

	return workflowRunCollection.Uniq(), nil
}

func (g *Github) WorkflowRuns(repo, branch string) (entity.WorkflowRunCollection, error) {
	response := struct { //nolint:exhaustruct
		TotalCount   int                  `json:"total_count"`
		WorkflowRuns []entity.WorkflowRun `json:"workflow_runs"`
	}{}

	err := g.client.Get(
		fmt.Sprintf("repos/%s/actions/runs?branch=%s", strings.Trim(repo, " /"), strings.Trim(branch, " ")),
		&response,
	)
	if err != nil {
		return entity.NewWorkflowRunCollection([]entity.WorkflowRun{}), err
	}

	return entity.NewWorkflowRunCollection(response.WorkflowRuns), nil
}
