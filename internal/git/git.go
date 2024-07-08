// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package git

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/cli/safeexec"
)

type git struct {
	path string
}

func newDefaultGit() (*git, error) {
	lookPath, err := safeexec.LookPath("git")
	if err != nil {
		return nil, fmt.Errorf("could not find git executable in PATH. error: %w", err)
	}

	return newGit(lookPath), nil
}

func newGit(path string) *git {
	return &git{path: path}
}

func (g *git) exec(args ...string) (stdOut, stdErr bytes.Buffer, err error) {
	return g.run(nil, args...)
}

func (g *git) run(env []string, args ...string) (stdOut, stdErr bytes.Buffer, err error) {
	cmd := exec.Command(g.path, args...)
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	if env != nil {
		cmd.Env = env
	}

	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("failed to run git: %s. error: %w", stdErr.String(), err)

		return
	}

	return
}
