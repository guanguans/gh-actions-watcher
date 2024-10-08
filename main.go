// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package main

import "github.com/guanguans/gh-actions-watcher/cmd"

// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
// https://cli.github.com/manual/gh_extension
// https://github.com/topics/gh-extension
// https://docs.github.com/zh/github-cli/github-cli/creating-github-cli-extensions
func main() {
	cmd.Execute()
}
