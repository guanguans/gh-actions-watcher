// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package git

import (
	"errors"
	"regexp"
	"strings"
)

type LocalGitRepo struct {
	git *git
}

func NewDefaultLocalGitRepo() (*LocalGitRepo, error) {
	defaultGit, err := newDefaultGit()
	if err != nil {
		return nil, err
	}

	return newLocalGitRepo(defaultGit), nil
}

func newLocalGitRepo(git *git) *LocalGitRepo {
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
	pattern := `(?:https:\/\/github\.com\/|git@github\.com:|git@github\.com:\/)([\w-]+\/[\w-]+)`
	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(githubRemoteURL)
	if len(match) > 1 {
		return match[1], nil
	}

	return "", errors.New("It seems you are executing this in a git repo that was not cloned from Github. Detected remote URL: `" + githubRemoteURL + "`")
}
