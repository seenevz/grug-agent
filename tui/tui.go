package tui

import (
	"fmt"
	"local-agent/utils"
)

type TUI struct {
	nextUserInput *utils.GetUserInput
}

func (t *TUI) PrintTool(name string, input any) {
	// \u001b[92m  bright green
	fmt.Printf("\u001b[92mTool\u001b[0m: %s(%s)\n", name, input)
}

func (t *TUI) PrintAgent(content string) {
	// \u001b[93m  bright yellow
	fmt.Printf("\u001b[93mGrug\u001b[0m: %s\n", content)
}

func (t *TUI) PrintMessage(message string) {
	fmt.Println(message)
}

func (t *TUI) PrintError(errorMsg string) {
	// \u001b[91m  bright red
	fmt.Printf("\u001b[91mError: %s\u001b[0m\n", errorMsg)
}

func (t *TUI) GetUserInput() (string, bool) {
	// \u001b[94m  bright blue
	fmt.Print("\u001b[94mYou\u001b[0m: ")
	return (*t.nextUserInput)()
}
func New() *TUI {
	t := TUI{nextUserInput: utils.ScanUserInput()}

	return &t
}
