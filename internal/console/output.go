// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package console

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/guanguans/gh-actions-watcher/internal/color"
)

type Output struct{}

func NewOutput() *Output {
	return &Output{}
}

func (o *Output) LineSuccess(message string) {
	o.Line(message, color.ColorGreen.String())
}

func (o *Output) LineWarning(message string) {
	o.Line(message, color.ColorOrange.String())
}

func (o *Output) LineError(message string) {
	o.Line(message, color.ColorRed.String())
}

func (o *Output) LineInfo(message string) {
	o.Line(message, color.ColorGray.String())
}

func (o *Output) Line(message, fg string) {
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

func (o *Output) BlockSuccess(message string) {
	o.Block(message, color.ColorGreen.String(), color.ColorWhite.String())
}

func (o *Output) BlockWarning(message string) {
	o.Block(message, color.ColorOrange.String(), color.ColorWhite.String())
}

func (o *Output) BlockError(message string) {
	o.Block(message, color.ColorRed.String(), color.ColorWhite.String())
}

func (o *Output) BlockInfo(message string) {
	o.Block(message, color.ColorGray.String(), color.ColorWhite.String())
}

func (o *Output) Block(message, bg, fg string) {
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
