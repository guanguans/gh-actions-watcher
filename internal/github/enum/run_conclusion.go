// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package enum

//go:generate go-enum --marshal --flag --names --values

// ENUM(action_required, cancelled, failure, neutral, success, skipped, stale, timed_out)
type RunConclusion string

func (x RunConclusion) HumanReadableValue() string {
	return humanReadableFor(x)
}

func (x RunConclusion) Color() string {
	switch x.String() {
	case RunConclusionActionRequired.String():
		return "#ff8c00" // orange
	case RunConclusionCancelled.String(), RunConclusionSkipped.String():
		return "#bfbfbf" // gray
	case RunConclusionFailure.String(), RunConclusionStale.String(), RunConclusionTimedOut.String():
		return "#ff0000" // red
	case RunConclusionNeutral.String(), RunConclusionSuccess.String():
		return "#008000" // green
	default:
		return "#bfbfbf" // gray
	}
}
