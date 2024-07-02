// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package entity

import "github.com/guanguans/gh-actions-watcher/internal/github/enum"

type WorkflowRun struct {
	Name       string `json:"name"`
	HtmlUrl    string `json:"html_url"`
	Status     string `json:"status"`
	Conclusion string `json:"conclusion"`
}

func (wr WorkflowRun) status() (enum.RunStatus, error) {
	return enum.ParseRunStatus(wr.Status)
}

func (wr WorkflowRun) GetListStatus() (enum.RunEnum, error) {
	if wr.Status == enum.RunStatusCompleted.String() {
		return wr.conclusion()
	}

	return wr.status()
}

func (wr WorkflowRun) conclusion() (enum.RunConclusion, error) {
	return enum.ParseRunConclusion(wr.Conclusion)
}

func (wr WorkflowRun) didComplete() bool {
	status, err := wr.status()
	if err != nil {
		return false
	}

	return status == enum.RunStatusCompleted
}

func (wr WorkflowRun) didNotComplete() bool {
	return !wr.didComplete()
}
