package game

import (
	"context"
	"strings"
	"time"
	"typeracer_cli/internal/config"
)

type Engine struct {
	config    *config.Config
	words     []string
	target    string
	startTime time.Time
	endTime   time.Time
}

type Result struct {
	WPM          float64
	Accuracy     float64
	TotalWords   int
	CorrectWords int
	Errors       int
	Duration     time.Duration
	Timestamp    time.Time
}

type DisplayUpdater interface {
	UpdateProgress(target, typed string, errors int)
}

type InputCapturer interface {
	StartCapture(ctx context.Context, inputChan chan<- rune)
}

func NewEngine(cfg *config.Config, words []string) *Engine {
	return &Engine{
		config: cfg,
		words:  words,
		target: strings.Join(words, " "),
	}
}

func (e *Engine) Start(display DisplayUpdater, input InputCapturer) *Result {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	inputChan := make(chan rune, 100)
	doneChan := make(chan bool, 1)

	go input.StartCapture(ctx, inputChan)

	if e.config.TimeLimit > 0 {
		go e.startTimer(ctx, time.Duration(e.config.TimeLimit)*time.Second, doneChan)
	}
	var typed strings.Builder
	var errors int
	var correctChars int

	e.startTime = time.Now()

	for {
		select {
		case char := <-inputChan:
			if char == '\n' || char == '\r' {
				doneChan <- true
				continue
			}

			typed.WriteRune(char)
			currentTyped := typed.String()

			if len(currentTyped) >= len(e.target) {
				doneChan <- true
				continue
			}

			display.UpdateProgress(e.target, currentTyped, errors)

			if len(currentTyped) > 0 {
				lastChar := currentTyped[len(currentTyped)-1]
				if len(currentTyped) <= len(e.target) {
					expectedChar := e.target[len(currentTyped)-1]
					if lastChar == expectedChar {
						correctChars++
					} else {
						errors++
					}
				}
			}

		case <-doneChan:
			e.endTime = time.Now()
			cancel()

			duration := e.endTime.Sub(e.startTime)
			typedText := typed.String()

			return e.calculateResult(typedText, errors, correctChars, duration)
		}
	}

}

func (e *Engine) startTimer(ctx context.Context, duration time.Duration, done chan<- bool) {
	timer := time.NewTimer(duration)
	defer timer.Stop()

	select {
	case <-timer.C:
		done <- true
	case <-ctx.Done():
		return
	}
}

func (e *Engine) calculateResult(typed string, errors, correctChars int, duration time.Duration) *Result {
	minutes := duration.Minutes()

	wpm := float64(correctChars) / 5.0 / minutes

	totalChars := len(typed)
	accuracy := 0.0
	if totalChars > 0 {
		accuracy = float64(correctChars) / float64(totalChars) * 100
	}

	wordsCompleted := 0
	typedWords := strings.Fields(typed)
	for i, word := range typedWords {
		if i < len(e.words) && word == e.words[i] {
			wordsCompleted++
		}
	}

	return &Result{
		WPM:          wpm,
		Accuracy:     accuracy,
		TotalWords:   len(e.words),
		CorrectWords: wordsCompleted,
		Errors:       errors,
		Duration:     duration,
		Timestamp:    time.Now(),
	}
}
