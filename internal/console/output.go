// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package console

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Output struct{}

func NewOutput() *Output {
	return &Output{}
}

func (o *Output) LineSuccess(message string) {
	o.Line(message, "#008000")
}

func (o *Output) LineWarning(message string) {
	o.Line(message, "#ff8c00")
}

func (o *Output) LineError(message string) {
	o.Line(message, "#ff0000")
}

func (o *Output) LineInfo(message string) {
	o.Line(message, "#bfbfbf")
}

func (o *Output) Line(message string, fg string) {
	fmt.Println(
		lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1).
			Foreground(lipgloss.Color(fg)).
			Width(120).
			Render(message),
	)
}

func (o *Output) NewLine(count int) {
	fmt.Print(strings.Repeat("\n", count))
}

func (o *Output) Success(message string) {
	o.Block(message, "#008000", "#ffffff")
}

func (o *Output) Warning(message string) {
	o.Block(message, "#ff8c00", "#ffffff")
}

func (o *Output) Error(message string) {
	o.Block(message, "#ff0000", "#ffffff")
}

func (o *Output) Info(message string) {
	o.Block(message, "#bfbfbf", "#ffffff")
}

func (o *Output) Block(message string, bg string, fg string) {
	fmt.Println(
		lipgloss.NewStyle().
			MarginTop(1).
			MarginBottom(1).
			Padding(1).
			Background(lipgloss.Color(bg)).
			Foreground(lipgloss.Color(fg)).
			Width(120).
			Render(message),
	)
}
