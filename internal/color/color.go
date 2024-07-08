// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package color

//go:generate go-enum --marshal --flag --names --values
//go:generate go-enum

// color ENUM(gray = "#bfbfbf", green = "#008000", orange = "#ff8c00", red = "#ff0000", white = "#ffffff", black = "#000000")
type color string
