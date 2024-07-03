// Copyright (c) 2024 guanguans<ityaozm@gmail.com>
// For the full copyright and license information, please view
// the LICENSE file that was distributed with this source code.
// https://github.com/guanguans/gh-actions-watcher

package console

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type consoleOutput struct{}

func NewConsoleOutput() *consoleOutput {
	return &consoleOutput{}
}

func (co *consoleOutput) lineSuccess(message string) {
	co.line(message, "#008000")
}

func (co *consoleOutput) lineWarning(message string) {
	co.line(message, "#ff8c00")
}

func (co *consoleOutput) lineError(message string) {
	co.line(message, "#ff0000")
}

func (co *consoleOutput) lineInfo(message string) {
	co.line(message, "#bfbfbf")
}

func (co *consoleOutput) line(message string, fg string) {
	style := lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
		Foreground(lipgloss.Color(fg)).
		Width(120)

	fmt.Println(style.Render(message))
}

func (co *consoleOutput) success(message string) {
	co.block(message, "#008000", "#ffffff")
}

func (co *consoleOutput) warning(message string) {
	co.block(message, "#ff8c00", "#ffffff")
}

func (co *consoleOutput) Error(message string) {
	co.block(message, "#ff0000", "#ffffff")
}

func (co *consoleOutput) info(message string) {
	co.block(message, "#bfbfbf", "#ffffff")
}

func (co *consoleOutput) block(message string, bg string, fg string) {
	style := lipgloss.NewStyle().
		MarginTop(1).
		MarginBottom(1).
		Padding(1).
		Background(lipgloss.Color(bg)).
		Foreground(lipgloss.Color(fg)).
		Width(120)

	fmt.Println(style.Render(message))
}

func (co *consoleOutput) newLine(count int) {
	for i := 0; i < count; i++ {
		fmt.Println("")
	}
}
