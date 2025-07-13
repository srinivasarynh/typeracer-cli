package ui

import (
	"bufio"
	"context"
	"os"
	"syscall"

	"golang.org/x/term"
)

type InputHandler struct {
	oldState *term.State
}

func NewInputHandler() *InputHandler {
	return &InputHandler{}
}

func (h *InputHandler) WaitForEnter() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func (h *InputHandler) StartCapture(ctx context.Context, InputChan chan<- rune) {
	oldState, err := term.MakeRaw(int(syscall.Stdin))
	if err != nil {
		return
	}
	h.oldState = oldState

	defer func() {
		if h.oldState != nil {
			term.Restore(int(syscall.Stdin), h.oldState)
		}
	}()

	buffer := make([]byte, 1)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			n, err := os.Stdin.Read(buffer)
			if err != nil || n == 0 {
				continue
			}
			char := rune(buffer[0])
			if char == 3 {
				return
			}
			InputChan <- char
		}
	}
}

func (h *InputHandler) Cleanup() {
	if h.oldState != nil {
		term.Restore(int(syscall.Stdin), h.oldState)
	}
}
