package tui

import (
	"fmt"
	"local-agent/utils"
	"sync"
	"time"
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

func (t *TUI) PrintAgentWaiting() func() {
	stop := make(chan bool)
	var needWait sync.WaitGroup

	cleanup := func() {
		fmt.Printf("\x1b[1K\r")
	}

	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		needWait.Add(1)

		fmt.Printf("Waiting for response ")

		for i := 0; ; i++ {
			if i > 3 {
				i = 0
				fmt.Printf("\x1b[1K\rWaiting for response ")
			}

			select {
			case <-ticker.C:
				fmt.Printf("+")
			case <-stop:
				cleanup()
				needWait.Done()
				return
			}
		}
	}()

	return func() { stop <- true; needWait.Wait() }
}

func (t *TUI) GetUserInput() (string, bool) {
	// \u001b[94m  bright blue
	fmt.Print("\u001b[94mYou\u001b[0m: ")
	str, ok := (*t.nextUserInput)()

	if !ok {
		return str, ok
	}

	// check if input has alphanumeric chars
	for _, r := range str {
		if r >= 48 && r <= 122 {
			return str, true
		}
	}

	return "", false
}

func New() *TUI {
	t := TUI{nextUserInput: utils.ScanUserInput()}

	return &t
}
