// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package enum

//go:generate go-enum --marshal --flag --names --values

// ENUM(queued, pending, in_progress, completed)
type RunStatus string

func (x RunStatus) HumanReadableValue() string {
	return humanReadableFor(x)
}

func (x RunStatus) Color() string {
	switch x.String() {
	case RunStatusQueued.String(), RunStatusPending.String():
		return "#bfbfbf" // gray
	case RunStatusInProgress.String():
		return "#ff8c00" // orange
	case RunStatusCompleted.String():
		return "#ff0000" // red
	default:
		return "#bfbfbf" // gray
	}
}
