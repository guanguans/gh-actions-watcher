// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package enum

import "github.com/guanguans/gh-actions-watcher/internal/color"

//go:generate go-enum --marshal --flag --names --values

// ENUM(action_required, cancelled, failure, neutral, success, skipped, stale, timed_out)
type RunConclusion string

func (x RunConclusion) HumanReadableValue() string {
	return humanReadableFor(x)
}

func (x RunConclusion) Color() string {
	switch x.String() {
	case RunConclusionActionRequired.String():
		return color.ColorOrange.String() // orange
	case RunConclusionCancelled.String(), RunConclusionSkipped.String():
		fallthrough // gray
	default:
		return color.ColorGray.String() // gray
	case RunConclusionFailure.String(), RunConclusionStale.String(), RunConclusionTimedOut.String():
		return color.ColorRed.String() // red
	case RunConclusionNeutral.String(), RunConclusionSuccess.String():
		return color.ColorGreen.String() // green
	}
}
