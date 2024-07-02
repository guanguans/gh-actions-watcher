// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package entity

import (
	"github.com/guanguans/gh-actions-watcher/internal/github/enum"
	"github.com/samber/lo"
)

type WorkflowRunCollection struct {
	workflowRuns []WorkflowRun
}

func NewWorkflowRunCollection(workflowRuns []WorkflowRun) WorkflowRunCollection {
	return WorkflowRunCollection{workflowRuns: workflowRuns}
}

func (wrc WorkflowRunCollection) ContainsActiveRuns() bool {
	return lo.ContainsBy(wrc.workflowRuns, func(workflowRun WorkflowRun) bool {
		return workflowRun.didNotComplete()
	})
}

func (wrc WorkflowRunCollection) AllCompletedSuccessfully() bool {
	if wrc.ContainsActiveRuns() {
		return false
	}

	return !lo.ContainsBy(wrc.workflowRuns, func(workflowRun WorkflowRun) bool {
		conclusion, err := workflowRun.conclusion()
		if err != nil {
			return false
		}

		return conclusion != enum.RunConclusionSuccess
	})
}

func (wrc WorkflowRunCollection) Uniq() WorkflowRunCollection {
	uniqWorkflowRuns := lo.UniqBy(wrc.workflowRuns, func(workflowRun WorkflowRun) string {
		return workflowRun.Name
	})

	return NewWorkflowRunCollection(uniqWorkflowRuns)
}

func (wrc WorkflowRunCollection) All() []WorkflowRun {
	return wrc.workflowRuns
}

func (wrc WorkflowRunCollection) IsEmpty() bool {
	return len(wrc.workflowRuns) == 0
}
