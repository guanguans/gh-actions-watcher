// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package enum

import "github.com/guanguans/gh-actions-watcher/internal/color"

//go:generate go-enum --marshal --flag --names --values
//go:generate go-enum

// RunStatus ENUM(queued, pending, in_progress, completed)
type RunStatus string

func (x RunStatus) HumanReadableValue() string {
	return humanReadableFor(x)
}

func (x RunStatus) Color() string {
	switch x.String() {
	case RunStatusInProgress.String():
		return color.ColorOrange.String() // orange
	case RunStatusCompleted.String():
		return color.ColorRed.String() // red
	case RunStatusQueued.String(), RunStatusPending.String():
		fallthrough // gray
	default:
		return color.ColorGray.String() // gray
	}
}
