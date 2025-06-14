package tui

import (
	"fmt"
	"local-agent/utils"
)

type TUI struct {
	nextUserInput *utils.GetUserInput
}

func (t *TUI) PrintTool(name string, input any) {
	fmt.Printf("\u001b[92mtool\u001b[0m: %s(%s)\n", name, input)
}

func (t *TUI) PrintAgent(content string) {
	fmt.Printf("\u001b[93mGrug\u001b[0m: %s\n", content)
}

func (t *TUI) PrintMessage(message string) {
	fmt.Println(message)
}

func (t *TUI) GetUserInput() (string, bool) {
	fmt.Print("\u001b[94mYou\u001b[0m: ")
	return (*t.nextUserInput)()
}
func New() *TUI {
	t := TUI{nextUserInput: utils.ScanUserInput()}

	return &t
}
