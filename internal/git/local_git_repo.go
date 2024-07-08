// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package git

import (
	"fmt"
	"regexp"
	"strings"
)

type LocalGitRepo struct {
	git *git
}

func NewDefaultLocalGitRepo() (*LocalGitRepo, error) {
	defaultGit, err := NewDefaultGit()
	if err != nil {
		return nil, err
	}

	return NewLocalGitRepo(defaultGit), nil
}

func NewLocalGitRepo(git *git) *LocalGitRepo {
	return &LocalGitRepo{git: git}
}

func (l LocalGitRepo) GetVendorAndRepo() (string, error) {
	githubRemoteURL, err := l.getConfiguredGitURL()
	if err != nil {
		return "", err
	}

	return l.extractVendorAndRepo(githubRemoteURL)
}

func (l LocalGitRepo) GetCurrentBranch() (string, error) {
	stdout, _, err := l.git.exec("rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(stdout.String()), nil
}

func (l LocalGitRepo) getConfiguredGitURL() (string, error) {
	stdout, _, err := l.git.exec("config", "--get", "remote.origin.url")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(stdout.String()), nil
}

func (l LocalGitRepo) extractVendorAndRepo(githubRemoteURL string) (string, error) {
	re := regexp.MustCompile(`(?:https://github\.com/|git@github\.com:|git@github\.com:/)([\w-]+/[\w-]+)`)
	match := re.FindStringSubmatch(githubRemoteURL)
	if len(match) < 1 {
		return "", fmt.Errorf(
			"it seems you are executing this in a git repo that was not cloned from Github. detected remote URL: `%s`",
			githubRemoteURL,
		)
	}

	return match[1], nil
}
