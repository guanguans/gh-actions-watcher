// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package enum

import (
	"fmt"
	"strings"
)

type RunEnum interface {
	HumanReadableValue() string
	Color() string
}

func humanReadableFor(str fmt.Stringer) string {
	return strings.ReplaceAll(strings.ToTitle(str.String()), "_", " ")
}
